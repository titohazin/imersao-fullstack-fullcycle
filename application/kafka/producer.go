package kafka

import (
	"fmt"
	"os"

	lkafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// NewKafkaProducer New Kafka Producer
func NewKafkaProducer() (*lkafka.Producer, error) {

	configMap := &lkafka.ConfigMap{
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
	}

	producer, err := lkafka.NewProducer(configMap)

	if err != nil {
		return nil, err
	}

	return producer, nil
}

// Publish Publish
func Publish(msg string, topic string, producer *lkafka.Producer, deliveryChannel chan lkafka.Event) error {

	message := &lkafka.Message{
		TopicPartition: lkafka.TopicPartition{Topic: &topic, Partition: lkafka.PartitionAny},
		Value:          []byte(msg),
	}

	err := producer.Produce(message, deliveryChannel)

	if err != nil {
		return err
	}

	return nil
}

// DeliveryReport Delivery Report
func DeliveryReport(deliveryChannel chan lkafka.Event) {

	for e := range deliveryChannel {

		switch ev := e.(type) {
		case *lkafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivery failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}

	}
}
