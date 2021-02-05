package kafka

import (
	"fmt"

	lkafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
)

// KafkaProcessor Kafka Processor
type KafkaProcessor struct {
	Database        *gorm.DB
	Producer        *lkafka.Producer
	DeliveryChannel chan lkafka.Event
}

// NewKafkaProcessor New Kafka Processor
func NewKafkaProcessor(database *gorm.DB, producer *lkafka.Producer, deliveryChannel chan lkafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Database:        database,
		Producer:        producer,
		DeliveryChannel: deliveryChannel,
	}
}

// Consume Consume
func (kafkaProc *KafkaProcessor) Consume() {

	configMap := &lkafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",
		"group.id":          "consumergroup",
		"auto.offset.reset": "earliest",
	}

	consumer, err := lkafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	topics := []string{"test"}
	consumer.SubscribeTopics(topics, nil)

	fmt.Println("Kfaka consumer has been created")

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			fmt.Println(string(msg.Value))
		}
	}
}
