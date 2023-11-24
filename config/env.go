package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvLoad(filenames ...string) {
	if len(filenames) == 0 {
		filenames = append(filenames, ".env")
	}
	for _, filename := range filenames {
		log.Printf("Loading configuration file: %s", filename)
		err := godotenv.Load(filename)
		if err != nil {
			log.Fatalf("Error loading configuration file: %s", filename)
		}
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
* XRPL protocol (compatible) server settings
 */
func EnvXrplWebsocketURL() string {
	return os.Getenv("XRPL_WEBSOCKET_URL")
}

func EnvXrplWebsocketFullHistoryURL() string {
	return os.Getenv("XRPL_WEBSOCKET_FULLHISTORY_URL")
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

func EnvEsNamespace() string {
	return os.Getenv("ELASTICSEARCH_NAMESPACE")
}

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
