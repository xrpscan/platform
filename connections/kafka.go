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

// Returns Kafka reader config with default settings applied
func NewReaderConfig() kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers: []string{config.EnvKafkaBootstrapServer()},
		GroupID: config.EnvKafkaGroupId(),
	}
}

var KafkaReaderLedger *kafka.Reader
var ledgerOnce sync.Once

// Create a new Kafka Reader connection to $namespace-platform-ledger topic
func NewLedgerReader() {
	ledgerOnce.Do(func() {
		cfg := NewReaderConfig()
		cfg.Topic = config.TopicLedgers()
		KafkaReaderLedger = kafka.NewReader(cfg)
	})
}

var KafkaReaderTransaction *kafka.Reader
var txOnce sync.Once

// Create a new Kafka Reader connection to $namespace-platform-transaction topic
func NewTransactionReader() {
	txOnce.Do(func() {
		cfg := NewReaderConfig()
		cfg.Topic = config.TopicTransactions()
		KafkaReaderTransaction = kafka.NewReader(cfg)
	})
}

var KafkaReaderValidation *kafka.Reader
var validationOnce sync.Once

// Create a new Kafka Reader connection to $namespace-platform-validation topic
func NewValidationReader() {
	validationOnce.Do(func() {
		cfg := NewReaderConfig()
		cfg.Topic = config.TopicValidations()
		KafkaReaderValidation = kafka.NewReader(cfg)
	})
}

var KafkaReaderPeerStatus *kafka.Reader
var peerstatusOnce sync.Once

// Create a new Kafka Reader connection to $namespace-platform-peerstatus topic
func NewPeerStatusReader() {
	peerstatusOnce.Do(func() {
		cfg := NewReaderConfig()
		cfg.Topic = config.TopicPeerStatus()
		KafkaReaderPeerStatus = kafka.NewReader(cfg)
	})
}

var KafkaReaderConsensus *kafka.Reader
var consensusOnce sync.Once

// Create a new Kafka Reader connection to $namespace-platform-consensus topic
func NewConsensusReader() {
	consensusOnce.Do(func() {
		cfg := NewReaderConfig()
		cfg.Topic = config.TopicConsensus()
		KafkaReaderConsensus = kafka.NewReader(cfg)
	})
}

var KafkaReaderPathFind *kafka.Reader
var pathfindOnce sync.Once

// Create a new Kafka Reader connection to $namespace-platform-pathfind topic
func NewPathFindReader() {
	pathfindOnce.Do(func() {
		cfg := NewReaderConfig()
		cfg.Topic = config.TopicPathFind()
		KafkaReaderPathFind = kafka.NewReader(cfg)
	})
}

var KafkaReaderManifest *kafka.Reader
var manifestOnce sync.Once

// Create a new Kafka Reader connection to $namespace-platform-manifest topic
func NewManifestReader() {
	manifestOnce.Do(func() {
		cfg := NewReaderConfig()
		cfg.Topic = config.TopicManifests()
		KafkaReaderManifest = kafka.NewReader(cfg)
	})
}

var KafkaReaderServer *kafka.Reader
var serverOnce sync.Once

// Create a new Kafka Reader connection to $namespace-platform-server topic
func NewServerReader() {
	serverOnce.Do(func() {
		cfg := NewReaderConfig()
		cfg.Topic = config.TopicServer()
		KafkaReaderServer = kafka.NewReader(cfg)
	})
}

var KafkaReaderDefault *kafka.Reader
var defaultOnce sync.Once

// Create a new Kafka Reader connection to $namespace-platform-default topic
func NewDefaultReader() {
	defaultOnce.Do(func() {
		cfg := NewReaderConfig()
		cfg.Topic = config.TopicDefault()
		KafkaReaderDefault = kafka.NewReader(cfg)
	})
}

func NewReaders() {
	NewLedgerReader()
	NewTransactionReader()
	NewValidationReader()
	NewPeerStatusReader()
	NewConsensusReader()
	NewPathFindReader()
	NewManifestReader()
	NewServerReader()
	NewDefaultReader()
}
