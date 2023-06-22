package producers

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/xrpl-go"
)

func SubscribeStreams() {
	connections.XrplClient.Subscribe([]string{
		xrpl.StreamTypeLedger,
		xrpl.StreamTypeTransaction,
		xrpl.StreamTypeValidations,
	})

	for {
		select {
		case ledger := <-connections.XrplClient.StreamLedger:
			ProduceLedger(connections.KafkaWriter, ledger)
			ProduceTx(connections.KafkaWriter, ledger)

		case validation := <-connections.XrplClient.StreamValidation:
			ProduceValidation(connections.KafkaWriter, validation)

		case transaction := <-connections.XrplClient.StreamTransaction:
			ProduceTransaction(connections.KafkaWriter, transaction)

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
