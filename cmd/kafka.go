package cmd

/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"os"

	lkafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/spf13/cobra"
	"github.com/titohazin/imersao-fullstack-fullcycle/application/kafka"
	"github.com/titohazin/imersao-fullstack-fullcycle/infra/db"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Start consuming transactions using Kafka",
	Run: func(cmd *cobra.Command, args []string) {

		database := db.ConnectDB(os.Getenv("env"))
		deliveryChannel := make(chan lkafka.Event)
		producer, err := kafka.NewKafkaProducer()

		if err != nil {
			panic(err)
		}

		kafka.Publish("Hello, Consumer!", "test", producer, deliveryChannel)
		go kafka.DeliveryReport(deliveryChannel)

		kafkaProcessor := kafka.NewKafkaProcessor(database, producer, deliveryChannel)
		kafkaProcessor.Consume()
	},
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kafkaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kafkaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
