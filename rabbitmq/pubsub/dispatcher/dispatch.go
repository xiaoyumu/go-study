package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/streadway/amqp"
)

func main() {

	log.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	log.Print("Starting RabbitMQ dispatcher ...")

	conn, err := amqp.Dial("amqp://godev:g0dev@192.168.1.51:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	log.Print("Connected to RabbitMQ successfully.")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	log.Print("Channel opened successfully.")

	watching(ch, "queue-pub-sub", func(message string, contentType string) {
		const targetURL string = "http://127.0.0.1:8923/system-event"

		resp, err := http.Post(targetURL,
			contentType,
			strings.NewReader(message))
		if err != nil {
			log.Printf("%v", err)
			return
		}
		defer resp.Body.Close()
		log.Printf("Response Code: %v", resp.StatusCode)
	})
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

func watching(ch *amqp.Channel, queueName string, dispatchAction func(message string, contentType string)) {

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
			log.Printf("Received a message: %s [%s]", d.MessageId, d.ContentType)
			dispatchAction(string(d.Body), d.ContentType)
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
