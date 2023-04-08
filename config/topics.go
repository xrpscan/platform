package config

import "fmt"

// Kafka topic for streaming xrpl.StreamTypeLedger messages
func TopicLedgers() string {
	return fmt.Sprintf("%s-platform-ledgers", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypeTransaction messages
func TopicTransactions() string {
	return fmt.Sprintf("%s-platform-transactions", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypeValidations messages
func TopicValidations() string {
	return fmt.Sprintf("%s-platform-validations", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypeManifests messages
func TopicManifests() string {
	return fmt.Sprintf("%s-platform-manifests", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypePeerStatus messages
func TopicPeerStatus() string {
	return fmt.Sprintf("%s-platform-peerstatus", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypeConsensus messages
func TopicConsensus() string {
	return fmt.Sprintf("%s-platform-consensus", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypeServer messages
func TopicServer() string {
	return fmt.Sprintf("%s-platform-server", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming messages that do not match any StreamType*
func TopicDefault() string {
	return fmt.Sprintf("%s-platform-default", EnvKafkaTopicNamespace())
}
