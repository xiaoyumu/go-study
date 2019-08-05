package main

import (
	"errors"
	"log"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

// KafkaSettings contains configuration settings for kafka connection
type KafkaSettings struct {
	conf  *kafka.ConfigMap
	topic *string
	debug bool
}

// NewSettings creates an instance of the KafkaSettings based on the given
// parameters.
func NewSettings(brokers *string, topic *string, debug bool) *KafkaSettings {
	return &KafkaSettings{
		conf: &kafka.ConfigMap{
			"session.timeout.ms": 6000,
			"bootstrap.servers":  *broker,
		},
		topic: topic,
		debug: debug,
	}
}

type kafkaDispatcher struct {
	settings *KafkaSettings
	producer *kafka.Producer
	closed bool
}

// NewKafkaDispatcher function creates an instance of Dispatcher
func NewKafkaDispatcher(settings *KafkaSettings) Dispatcher {
	return &kafkaDispatcher{
		settings: settings,
	}
}

func (kd *kafkaDispatcher) Connect() error {
	log.Println("Connecting to kafka broker ...")
	p, err := kafka.NewProducer(kd.settings.conf)
	if err != nil {
		return err
	}

	kd.producer = p
	go func() {
		// Subscribe events channel to check delivery result.
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					kd.debug("Delivery failed: %v Error: %s",
						ev.TopicPartition,
						ev.TopicPartition.Error)
				} else {
					kd.debug("Delivered message to %v",
						ev.TopicPartition)
				}
			}
		}
	}()

	log.Println("kafka producer has been initialized.")
	return nil
}

func (kd *kafkaDispatcher) debug(format string, args ...interface{}) {
	if !kd.settings.debug {
		return
	}
	log.Printf(format, args...)
}

func (kd *kafkaDispatcher) DispatchMessage(msg *Message) error {
	return kd.Dispatch(msg.Key, msg.Value)
}

func (kd *kafkaDispatcher) Dispatch(key []byte, value []byte) error{
	if kd.producer == nil {
		return errors.New("kafka producer was not initialized")
	}

	kd.debug("Dispatching message ...")

	return kd.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     kd.settings.topic,
			Partition: kafka.PartitionAny,
		},
		Key:   key,
		Value: value,
	}, nil)
}

func (kd *kafkaDispatcher) Close() {
	if kd.closed {
		return
	}

	kd.closed = true
	// Wait for 5 seconds for message deliveries before shutdown.
	kd.producer.Flush(5 * 1000)
	kd.producer.Close()
}
