package consumers

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/processor"
)

func RunTransactionConsumer() {
	c := connections.KafkaConsumer
	err := c.SubscribeTopics([]string{"test.messages"}, nil)
	if err != nil {
		panic("Cannot subscribe to topic")
	}

	run := true
	for run {
		ev := c.Poll(100)
		if ev == nil {
			fmt.Println("Poll result: nil")
			continue
		}

		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("Message on %s: (%s)\n", e.TopicPartition, string(e.Value))
			processor.IndexTransaction(string(e.Value))
		case kafka.Error:
			fmt.Printf("Error: %v: %v\n", e.Code(), e)
		}
	}
}
