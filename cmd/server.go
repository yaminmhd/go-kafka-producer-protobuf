package cmd

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
	"github.com/yaminmhd/go-kafka-producer-protobuf/config"
	"github.com/yaminmhd/go-kafka-producer-protobuf/generator"
	"log"
)

func newSyncProducer(brokerList []string) (sarama.SyncProducer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Return.Successes = true
	// TODO additional producer configs

	producer, err := sarama.NewSyncProducer(brokerList, saramaConfig)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

func ProduceMessage(){
	topic := config.Topic()
	brokerList := config.Brokers()

	producer, err := newSyncProducer(brokerList)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	defer func() {
		if err := producer.Close(); err != nil{
			panic(err)
		}
	}()

	personMessage := generator.NewPerson()
	messageBytes, err := proto.Marshal(personMessage)
	if err != nil {
		log.Fatalln("Failed to marshal person proto:", err)
	}

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(messageBytes),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	fmt.Printf("Person message sent: %s", personMessage)
}
