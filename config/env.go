package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

/*
* Service settings
 */

// Get default log level
func EnvLogLevel() string {
	return os.Getenv("LOG_LEVEL")
}

// Get default log type
func EnvLogType() string {
	return os.Getenv("LOG_TYPE")
}

/*
* Rippled settings
 */
func EnvRippledURL() string {
	return os.Getenv("RIPPLED_URL")
}

func EnvRippledFullHistoryURL() string {
	return os.Getenv("RIPPLED_FULLHISTORY_URL")
}

/*
* Kafka settings
 */
func EnvKafkaBootstrapServer() string {
	return os.Getenv("KAFKA_BOOTSTRAP_SERVER")
}

func EnvKafkaGroupId() string {
	return os.Getenv("KAFKA_GROUP_ID")
}

func EnvKafkaTopicNamespace() string {
	return os.Getenv("KAFKA_TOPIC_NAMESPACE")
}

/*
* Elasticsearch settings
 */

func EnvEsURL() string {
	return os.Getenv("ELASTICSEARCH_URL")
}

func EnvEsUsername() string {
	return os.Getenv("ELASTICSEARCH_USERNAME")
}

func EnvEsPassword() string {
	return os.Getenv("ELASTICSEARCH_PASSWORD")
}

func EnvEsFingerprint() string {
	return os.Getenv("ELASTICSEARCH_FINGERPRINT")
}
