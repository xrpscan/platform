package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
	"github.com/xrpscan/platform/producers"
)

const defaultIndexFrom int = 82000000
const defaultIndexTo int = 82001000

var indexFrom int
var indexTo int
var configFile string

func main() {
	flag.IntVar(&indexFrom, "from", defaultIndexFrom, "From ledger index")
	flag.IntVar(&indexTo, "to", defaultIndexTo, "To ledger index")
	flag.StringVar(&configFile, "config", ".env", "Environment config file")
	flag.Parse()

	fmt.Printf("Using environment config file %s\n", configFile)
	config.EnvLoad(configFile)

	// Ledgers are backfilled in chronological order. Therefore, --from ledger
	// index must be less than --to ledger index.
	if indexFrom > indexTo {
		log.Fatalf("From ledger (%d) must be less than To ledger (%d)\n",
			indexFrom,
			indexTo)
	}

	// Initialize connections to services
	logger.New()
	connections.NewWriter()
	connections.NewXrplClient()

	// Fetch ledger and queue transactions for indexing
	for ledgerIndex := indexFrom; ledgerIndex <= indexTo; ledgerIndex++ {
		backfillLedger(ledgerIndex)
	}

	connections.CloseWriter()
	connections.CloseXrplClient()
}

func backfillLedger(ledgerIndex int) {
	log.Println("Backfilling ledger:", ledgerIndex)
	ledger := models.LedgerStream{
		Type:        models.LEDGER_STREAM_TYPE,
		LedgerIndex: uint32(ledgerIndex),
	}
	ledgerJSON, _ := json.Marshal(ledger)
	producers.ProduceTransactions(connections.KafkaWriter, ledgerJSON)
}
