package config

import "fmt"

const serviceName = "platform"

// Kafka topic for streaming xrpl.StreamTypeLedger messages
func TopicLedgers() string {
	return fmt.Sprintf("%s-%s-ledgers", EnvKafkaTopicNamespace(), serviceName)
}

// Kafka topic for streaming xrpl.StreamTypeTransaction messages
func TopicTransactions() string {
	return fmt.Sprintf("%s-%s-transactions", EnvKafkaTopicNamespace(), serviceName)
}

// Kafka topic for streaming xrpl.StreamTypeValidations messages
func TopicValidations() string {
	return fmt.Sprintf("%s-%s-validations", EnvKafkaTopicNamespace(), serviceName)
}

// Kafka topic for streaming xrpl.StreamTypeManifests messages
func TopicManifests() string {
	return fmt.Sprintf("%s-%s-manifests", EnvKafkaTopicNamespace(), serviceName)
}

// Kafka topic for streaming xrpl.StreamTypePeerStatus messages
func TopicPeerStatus() string {
	return fmt.Sprintf("%s-%s-peerstatus", EnvKafkaTopicNamespace(), serviceName)
}

// Kafka topic for streaming xrpl.StreamTypeConsensus messages
func TopicConsensus() string {
	return fmt.Sprintf("%s-%s-consensus", EnvKafkaTopicNamespace(), serviceName)
}

// Kafka topic for streaming xrpl.StreamTypePathFind messages
func TopicPathFind() string {
	return fmt.Sprintf("%s-%s-pathfind", EnvKafkaTopicNamespace(), serviceName)
}

// Kafka topic for streaming xrpl.StreamTypeServer messages
func TopicServer() string {
	return fmt.Sprintf("%s-%s-server", EnvKafkaTopicNamespace(), serviceName)
}

// Kafka topic for streaming messages that do not match any StreamType*
func TopicDefault() string {
	return fmt.Sprintf("%s-%s-default", EnvKafkaTopicNamespace(), serviceName)
}
