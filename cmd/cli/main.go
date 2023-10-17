package main

import (
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
	"github.com/xrpscan/platform/producers"
)

const defaultIndexFrom int = 82000000
const defaultIndexTo int = 82001000
const defaultMinDelay int64 = 100 // milliseconds

var fIndexFrom int
var fIndexTo int
var fConfigFile string
var fVerbose bool
var fXrplServer string
var fMinDelay int64

func setFlags() {
	flag.IntVar(&fIndexFrom, "from", defaultIndexFrom, "From ledger index")
	flag.IntVar(&fIndexTo, "to", defaultIndexTo, "To ledger index")
	flag.StringVar(&fConfigFile, "config", ".env", "Environment config file")
	flag.BoolVar(&fVerbose, "verbose", false, "Make the command more talkative")
	flag.StringVar(&fXrplServer, "server", "", "XRPL protocol compatible server to connect")
	flag.Int64Var(&fMinDelay, "delay", defaultMinDelay, "Minimum delay (ms) between requests to XRPL server")
	flag.Parse()
}

func clog(message ...string) {
	if fVerbose {
		log.Println(message)
	}
}

func main() {
	setFlags()
	clog("Using environment config file: ", fConfigFile)
	config.EnvLoad(fConfigFile)

	// Ledgers are backfilled in chronological order. Therefore, --from ledger
	// index must be less than --to ledger index.
	if fIndexFrom > fIndexTo {
		log.Fatalf("From ledger (%d) must be less than To ledger (%d)\n",
			fIndexFrom,
			fIndexTo)
	}

	// If websocket url is not provided in the cli, use the url from environment
	wsURL := fXrplServer
	if wsURL == "" {
		wsURL = config.EnvXrplWebsocketURL()
	}

	// Initialize connections to services
	logger.New()
	connections.NewWriter()
	connections.NewXrplClientWithURL(wsURL)
	defer connections.CloseWriter()
	defer connections.CloseXrplClient()

	// Fetch ledger and queue transactions for indexing
	for ledgerIndex := fIndexFrom; ledgerIndex <= fIndexTo; ledgerIndex++ {
		startTime := time.Now().UnixNano() / int64(time.Millisecond)
		backfillLedger(ledgerIndex)
		reqDuration := time.Now().UnixNano()/int64(time.Millisecond) - startTime

		// Honor fair usage policy and wait before sending next request
		delayRequired := fMinDelay - reqDuration
		if delayRequired > 0 {
			time.Sleep(time.Duration(delayRequired) * time.Millisecond)
		}
	}
}

func backfillLedger(ledgerIndex int) {
	log.Println("Backfilling ledger:", ledgerIndex)
	ledger := models.LedgerStream{
		Type:        models.LEDGER_STREAM_TYPE,
		LedgerIndex: uint32(ledgerIndex),
	}
	ledgerJSON, _ := json.Marshal(ledger)
	producers.ProduceLedger(connections.KafkaWriter, ledgerJSON)
	backfillTransactions(ledgerJSON)
}

func backfillTransactions(ledgerJSON []byte) {
	producers.ProduceTransactions(connections.KafkaWriter, ledgerJSON)
}
