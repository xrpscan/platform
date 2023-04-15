package consumers

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/processor"
)

func RunConsumer(conn *kafka.Reader, callback func(m kafka.Message)) {
	ctx := context.Background()
	for {
		m, err := conn.FetchMessage(ctx)
		if err != nil {
			break
		}
		callback(m)

		// Commit message
		if err := conn.CommitMessages(ctx, m); err != nil {
			log.Println("Failed to commit messages to kafka: ", err)
		}
	}
}

func RunConsumers() {
	go RunConsumer(connections.KafkaReaderLedger, processor.IndexLedger)
	go RunConsumer(connections.KafkaReaderTransaction, processor.IndexTransaction)
	go RunConsumer(connections.KafkaReaderValidation, processor.IndexValidation)
	go RunConsumer(connections.KafkaReaderDefault, processor.PrintMessage)
}
