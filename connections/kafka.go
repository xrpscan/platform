package connections

import (
	"sync"

	"github.com/segmentio/kafka-go"
	"github.com/xrpscan/platform/config"
)

var KafkaWriter *kafka.Writer
var wOnce sync.Once

// Common Kafka writer for the application. Every message must
// specify the Topic where it must be written to.
func NewWriter() {
	wOnce.Do(func() {
		KafkaWriter = &kafka.Writer{
			Addr:     kafka.TCP(config.EnvKafkaBootstrapServer()),
			Balancer: &kafka.LeastBytes{},
			Async:    true,
		}
	})
}

var KafkaReaderLedger *kafka.Reader
var lOnce sync.Once

func NewLedgerReader() {
	lOnce.Do(func() {
		KafkaReaderLedger = kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.EnvKafkaBootstrapServer()},
			GroupID: config.EnvKafkaGroupId(),
			Topic:   config.TopicLedgers(),
		})
	})
}

var KafkaReaderTransaction *kafka.Reader
var krtOnce sync.Once

func NewTransactionReader() {
	krtOnce.Do(func() {
		KafkaReaderTransaction = kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.EnvKafkaBootstrapServer()},
			GroupID: config.EnvKafkaGroupId(),
			Topic:   config.TopicTransactions(),
		})
	})
}

var KafkaReaderValidation *kafka.Reader
var krvOnce sync.Once

func NewValidationReader() {
	krvOnce.Do(func() {
		KafkaReaderValidation = kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.EnvKafkaBootstrapServer()},
			GroupID: config.EnvKafkaGroupId(),
			Topic:   config.TopicValidations(),
		})
	})
}

var KafkaReaderDefault *kafka.Reader
var krdOnce sync.Once

func NewDefaultReader() {
	krdOnce.Do(func() {
		KafkaReaderDefault = kafka.NewReader(kafka.ReaderConfig{
			Brokers: []string{config.EnvKafkaBootstrapServer()},
			GroupID: config.EnvKafkaGroupId(),
			Topic:   config.TopicDefault(),
		})
	})
}

func NewReaders() {
	NewLedgerReader()
	NewTransactionReader()
	NewValidationReader()
	NewDefaultReader()
}
