package kafka

import (
	"fmt"
	"os"

	lkafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	"github.com/titohazin/imersao-fullstack-fullcycle/application/factory"
	appTransactionModel "github.com/titohazin/imersao-fullstack-fullcycle/application/model"
	"github.com/titohazin/imersao-fullstack-fullcycle/application/usecase"
	"github.com/titohazin/imersao-fullstack-fullcycle/domain/model"
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
		"bootstrap.servers": os.Getenv("kafkaBootstrapServers"),
		"group.id":          os.Getenv("kafkaConsumerGroupId"),
		"auto.offset.reset": "earliest",
	}

	consumer, err := lkafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	topics := []string{os.Getenv("kafkaTransactionTopic"), os.Getenv("kafkaTransactionConfirmationTopic")}
	consumer.SubscribeTopics(topics, nil)

	fmt.Println("Kfaka consumer has been created")

	for {
		msg, err := consumer.ReadMessage(-1)

		if err == nil {
			fmt.Println(string(msg.Value))
		}
	}
}

func (kafkaProc *KafkaProcessor) processMessage(msg *lkafka.Message) {
	transactionsTopic := "transactions"
	transactionConfirmationTopic := "transaction_confirmation"

	switch topic := *msg.TopicPartition.Topic; topic {
	case transactionsTopic:
		kafkaProc.processTransaction(msg)
	case transactionConfirmationTopic:
		kafkaProc.processTransactionConfirmation(msg)
	default:
		fmt.Println("is not a valid topic", string(msg.Value))
	}
}

func (kafkaProc *KafkaProcessor) processTransaction(msg *lkafka.Message) error {

	transaction := appTransactionModel.NewTransaction()
	err := transaction.JSONToModel(msg.Value)

	if err != nil {
		return err
	}

	transactionUseCase := factory.TranactionUseCaseFactory(kafkaProc.Database)

	createdTransaction, err := transactionUseCase.Register(
		transaction.AccountID,
		transaction.Amount,
		transaction.PixKeyTo,
		transaction.PixKeyToKind,
		transaction.Description,
	)

	if err != nil {
		fmt.Println("erro registering transaction", err)
		return err
	}

	transaction.ID = createdTransaction.ID
	transactionJSON, err := transaction.ModelToJSON()

	if err != nil {
		return err
	}

	topic := "bank" + createdTransaction.PixKeyTo.Account.Bank.Code
	err = Publish(string(transactionJSON), topic, kafkaProc.Producer, kafkaProc.DeliveryChannel)

	if err != nil {
		return err
	}

	return nil
}

func (kafkaProc *KafkaProcessor) processTransactionConfirmation(msg *lkafka.Message) error {
	transaction := appTransactionModel.NewTransaction()
	err := transaction.JSONToModel(msg.Value)

	if err != nil {
		return err
	}

	transactionUseCase := factory.TranactionUseCaseFactory(kafkaProc.Database)

	if transaction.Status == model.TransactionConfirmed {

		err = kafkaProc.confirmTransaction(transaction, &transactionUseCase)

		if err != nil {
			return err
		}
	} else if transaction.Status == model.TransactionCompleted {

		_, err := transactionUseCase.Complete(transaction.ID)

		if err != nil {
			return err
		}
	}

	return nil
}

func (kafkaProc *KafkaProcessor) confirmTransaction(transaction *appTransactionModel.Transaction, transactionUseCase *usecase.TransactionUseCase) error {

	confirmTransaction, err := transactionUseCase.Confirm(transaction.ID)

	if err != nil {
		return err
	}

	transactionJSON, err := transaction.ModelToJSON()

	if err != nil {
		return err
	}

	topic := "bank" + confirmTransaction.AccountFrom.Bank.Code
	err = Publish(string(transactionJSON), topic, kafkaProc.Producer, kafkaProc.DeliveryChannel)

	if err != nil {
		return err
	}

	return nil
}
