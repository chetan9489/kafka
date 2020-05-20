package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.Retry.Max = 5

	// brokers := []string{"192.168.59.103:9092"}
	brokers := []string{"localhost:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		// Should not reach here
		fmt.Println("Step1-Panic")
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			// Should not reach here
			fmt.Println("Step2-Panic")
			panic(err)
		}
	}()

	topic := "test"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("Something Cool"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Step3-Panic")
		panic(err)
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
