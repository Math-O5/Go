package kafka

import (
	ckafka "github.cm/confluentine/confluent-kafka-go/kafka"
)

func NewKafKaProducer() *ckafka.Producer {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers":"kafka:9092",
	}

	p, err := ckafka.NewProducer(configMap)

	if err != nil {
		panic(err)
	}

	return p
}

func Publiher(msg string, topic string, producer *ckafka.Producer, deliveryChan chan ckafka.Event) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: &topic, Partition: ckafka.PartitionAny}
		Value: []byte(msg),
	}

	err := producer.Produce(message, deliveryChan)

	if err != nil {
		return err
	}

	return nil
}

/* This is a loop: waiting message */
func DeliveryReport(deliveryChan chan ckafka.Event) {
	for e := range deliveryChan {
		switch ev := e.(type) {
		case *ckafka.Messafe:
			if ev.TopicPartition.Error != nil {
				fmt.Println("Delivere failed:", ev.TopicPartition)
			} else {
				fmt.Println("Delivered message to:", ev.TopicPartition)
			}
		}

	}
}