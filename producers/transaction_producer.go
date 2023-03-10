package producers

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

func Produce(p *kafka.Producer, message string) {
	topic := "test.messages"
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
		Key:            []byte("42"),
	}, nil)
}
