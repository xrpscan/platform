package config

import "fmt"

// Kafka topic for streaming xrpl.StreamTypeLedger messages
func TopicLedgers() string {
	return fmt.Sprintf("%s-ledgers", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypeTransaction messages
func TopicTransactions() string {
	return fmt.Sprintf("%s-transactions", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypeValidations messages
func TopicValidations() string {
	return fmt.Sprintf("%s-validations", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypeManifests messages
func TopicManifests() string {
	return fmt.Sprintf("%s-manifests", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypePeerStatus messages
func TopicPeerStatus() string {
	return fmt.Sprintf("%s-peerstatus", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypeConsensus messages
func TopicConsensus() string {
	return fmt.Sprintf("%s-consensus", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypePathFind messages
func TopicPathFind() string {
	return fmt.Sprintf("%s-pathfind", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming xrpl.StreamTypeServer messages
func TopicServer() string {
	return fmt.Sprintf("%s-server", EnvKafkaTopicNamespace())
}

// Kafka topic for streaming messages that do not match any StreamType*
func TopicDefault() string {
	return fmt.Sprintf("%s-default", EnvKafkaTopicNamespace())
}
