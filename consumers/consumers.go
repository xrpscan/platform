package consumers

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/indexer"
	"github.com/xrpscan/platform/logger"
)

// Serial consumer (based on callback function) for low volume message streams
func RunConsumer(conn *kafka.Reader, callback func(m kafka.Message)) {
	ctx := context.Background()
	for {
		m, err := conn.FetchMessage(ctx)
		if err != nil {
			break
		}
		callback(m)

		if err := conn.CommitMessages(ctx, m); err != nil {
			logger.Log.Error().Err(err).Msg("Failed to commit kafka message")
		}
	}
}

// Bulk message consumer (based on channel) for high volume message streams
func RunBulkConsumer(conn *kafka.Reader, callback func(<-chan kafka.Message)) {
	ctx := context.Background()
	ch := make(chan kafka.Message)
	go callback(ch)

	for {
		m, err := conn.FetchMessage(ctx)
		if err != nil {
			break
		}

		ch <- m

		if err := conn.CommitMessages(ctx, m); err != nil {
			logger.Log.Error().Err(err).Msg("Failed to commit kafka message")
		}
	}
}

// Run all consumers
func RunConsumers() {
	go RunBulkConsumer(connections.KafkaReaderLedger, indexer.BulkIndexLedger)
	go RunBulkConsumer(connections.KafkaReaderValidation, indexer.BulkIndexValidation)
	go RunBulkConsumer(connections.KafkaReaderTransaction, indexer.BulkIndexTransaction)

	go RunConsumer(connections.KafkaReaderPeerStatus, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderConsensus, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderPathFind, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderManifest, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderServer, indexer.PrintMessage)
	go RunConsumer(connections.KafkaReaderDefault, indexer.PrintMessage)
}
