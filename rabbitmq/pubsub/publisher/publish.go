package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
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

	for x := 0; x < 100; x++ {

		event := SystemEvent{
			ID:        fmt.Sprint(uuid.New()),
			Message:   fmt.Sprintf("This is auto-message %v", x),
			EventTime: time.Now(),
		}
		jsonContent, serializationError := json.Marshal(event)
		if serializationError != nil {
			log.Fatal(serializationError)
		}
		sending(ch, "queue-pub-sub", string(jsonContent), "application/json")
	}

}

// The SystemEvent represents an event
type SystemEvent struct {
	ID        string    `json:"id"`
	Message   string    `json:"message"`
	EventTime time.Time `json:"eventTime"`
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

func sending(ch *amqp.Channel, queueName string, body string, contentType string) {

	// To send, we must declare a queue for us to send to; then we can publish a message to the queue:
	q := declareQueue(ch, queueName)

	log.Printf("Publishing message to queue [%s] ...", queueName)
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        []byte(body), // Convert body string to []byte
		})
	failOnError(err, "Failed to publish a message")

	log.Printf("Message has been sent to queue [%s] successfully.", queueName)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
