package connections

import (
	"log"
)

func CloseWriter() {
	if KafkaWriter != nil {
		err := KafkaWriter.Close()
		if err != nil {
			log.Println("Error closing Kafka writer connection")
		}
	}
}

func CloseReaders() {
	if KafkaReaderLedger != nil {
		err := KafkaReaderLedger.Close()
		if err != nil {
			log.Println("Error closing Kafka Ledger reader connection")
		}
	}

	if KafkaReaderTransaction != nil {
		err := KafkaReaderTransaction.Close()
		if err != nil {
			log.Println("Error closing Kafka Transaction reader connection")
		}
	}

	if KafkaReaderValidation != nil {
		err := KafkaReaderValidation.Close()
		if err != nil {
			log.Println("Error closing Kafka Validation reader connection")
		}
	}

	if KafkaReaderPeerStatus != nil {
		err := KafkaReaderPeerStatus.Close()
		if err != nil {
			log.Println("Error closing Kafka PeerStatus reader connection")
		}
	}

	if KafkaReaderConsensus != nil {
		err := KafkaReaderConsensus.Close()
		if err != nil {
			log.Println("Error closing Kafka Consensus reader connection")
		}
	}

	if KafkaReaderPathFind != nil {
		err := KafkaReaderPathFind.Close()
		if err != nil {
			log.Println("Error closing Kafka PathFind reader connection")
		}
	}

	if KafkaReaderManifest != nil {
		err := KafkaReaderManifest.Close()
		if err != nil {
			log.Println("Error closing Kafka Manifest reader connection")
		}
	}

	if KafkaReaderServer != nil {
		err := KafkaReaderServer.Close()
		if err != nil {
			log.Println("Error closing Kafka Server reader connection")
		}
	}

	if KafkaReaderDefault != nil {
		err := KafkaReaderDefault.Close()
		if err != nil {
			log.Println("Error closing Kafka Default reader connection")
		}
	}
}

func CloseEsClient() {
}

func CloseXrplClient() {
	if XrplClient != nil {
		err := XrplClient.Close()
		if err != nil {
			log.Println("Error closing xrpl connection")
		}
	}
}

func CloseAll() {
	log.Println("Closing all connections")
	CloseWriter()
	CloseReaders()
	CloseEsClient()
	CloseXrplClient()
}
