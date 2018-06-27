package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://godev:g0dev@192.168.1.51:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	log.Print("Connected to RabbitMQ successfully.")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	log.Print("Channel opened successfully.")

	//sending(ch, "hello", "world3")
	receiving(ch, "hello")
}

func declareQueue(ch *amqp.Channel, queueName string) *amqp.Queue {

	log.Printf("Declaring Queue [%s] ...", queueName)
	// To send, we must declare a queue for us to send to; then we can publish a message to the queue:
	q, err := ch.QueueDeclare(
		"hello", // Queue name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	log.Printf("Queue [%s] declared successfully.", queueName)

	return &q
}

func sending(ch *amqp.Channel, queueName string, body string) {

	// To send, we must declare a queue for us to send to; then we can publish a message to the queue:
	q := declareQueue(ch, queueName)

	log.Printf("Publishing message to queue [%s] ...", queueName)
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body), // Convert body string to []byte
		})
	failOnError(err, "Failed to publish a message")

	log.Printf("Message has been sent to queue [%s] successfully.", queueName)
	log.Printf("Message Body:")
	log.Print(body)
}

func receiving(ch *amqp.Channel, queueName string) {

	q := declareQueue(ch, queueName)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s [%s]", d.Body, d.ContentType)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
