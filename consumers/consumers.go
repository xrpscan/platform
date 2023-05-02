package consumers

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/indexer"
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
	go RunConsumer(connections.KafkaReaderLedger, indexer.Test)
	go RunConsumer(connections.KafkaReaderTransaction, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderValidation, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderPeerStatus, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderConsensus, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderPathFind, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderManifest, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderServer, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderDefault, indexer.PrintMessage)
}
