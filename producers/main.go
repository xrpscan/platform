package producers

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/xrpl-go"
)

/*
* TLDR - Do not subscribe to xrpl.StreamTypeTransactions
*
* XRPL `transaction` stream messages are incompatible with rippled's native
* transaction format. Therefore, this service does not process transactions
* streamed on `xrpl.StreamTypeTransactions` stream. Instead, we listen to the
* ledger stream, fetch transactions from rippled, and add those transactions to
* the Kafka topic.
 */
func SubscribeStreams() {
	response, err := connections.XrplClient.Subscribe([]string{
		xrpl.StreamTypeLedger,
		xrpl.StreamTypeValidations,
	})
	if err != nil {
		logger.Log.Error().Err(err).Msg("xrpl.Subscribe")
	}
	if response["status"].(string) == "error" {
		logger.Log.Error().Any("error", response["error"]).Any("id", response["id"]).Any("error_message", response["error_message"]).Msg("xrpl.Subscribe")
	} else {
		logger.Log.Debug().Any("status", response["status"]).Any("id", response["id"]).Any("result", response["result"]).Msg("xrpl.Subscribe")
	}

	for {
		select {
		case ledger := <-connections.XrplClient.StreamLedger:
			ProduceLedger(connections.KafkaWriter, ledger)
			ProduceTransactions(connections.KafkaWriter, ledger)

		case validation := <-connections.XrplClient.StreamValidation:
			ProduceValidation(connections.KafkaWriter, validation)

		case peerStatus := <-connections.XrplClient.StreamPeerStatus:
			Produce(connections.KafkaWriter, peerStatus, config.TopicPeerStatus())

		case consensus := <-connections.XrplClient.StreamConsensus:
			Produce(connections.KafkaWriter, consensus, config.TopicConsensus())

		case pathFind := <-connections.XrplClient.StreamPathFind:
			Produce(connections.KafkaWriter, pathFind, config.TopicPathFind())

		case manifest := <-connections.XrplClient.StreamManifest:
			Produce(connections.KafkaWriter, manifest, config.TopicManifests())

		case server := <-connections.XrplClient.StreamServer:
			Produce(connections.KafkaWriter, server, config.TopicServer())

		case defaultObject := <-connections.XrplClient.StreamDefault:
			fmt.Println(string(defaultObject))
			Produce(connections.KafkaWriter, defaultObject, config.TopicDefault())
		}
	}
}

func Produce(w *kafka.Writer, message []byte, topic string) {
	messageKey := uuid.NewString()
	if topic == "" {
		topic = config.TopicDefault()
	}

	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Topic: topic,
			Key:   []byte(messageKey),
			Value: message,
		},
	)
	if err != nil {
		log.Printf("Failed to write message: %s", err)
	}
}
