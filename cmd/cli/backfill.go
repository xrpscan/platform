/**
* This file implements `platform-cli backfill` subcommand
 */

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/xrpscan/platform/config"
	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/platform/logger"
	"github.com/xrpscan/platform/models"
	"github.com/xrpscan/platform/producers"
	"github.com/xrpscan/platform/signals"
)

const BackfillCommandName = "backfill"
const defaultIndexFrom int = 82000000
const defaultIndexTo int = 82001000
const defaultMinDelay int64 = 100 // milliseconds

type BackfillCommand struct {
	fs          *flag.FlagSet
	fConfigFile string
	fXrplServer string
	fIndexFrom  int
	fIndexTo    int
	fMinDelay   int64
	fVerbose    bool
}

func NewBackfillCommand() *BackfillCommand {
	cmd := &BackfillCommand{
		fs: flag.NewFlagSet(BackfillCommandName, flag.ExitOnError),
	}

	cmd.fs.IntVar(&cmd.fIndexFrom, "from", defaultIndexFrom, "From ledger index")
	cmd.fs.IntVar(&cmd.fIndexTo, "to", defaultIndexTo, "To ledger index")
	cmd.fs.StringVar(&cmd.fConfigFile, "config", ".env", "Environment config file")
	cmd.fs.BoolVar(&cmd.fVerbose, "verbose", false, "Make the command more talkative")
	cmd.fs.StringVar(&cmd.fXrplServer, "server", "", "XRPL protocol compatible server to connect")
	cmd.fs.Int64Var(&cmd.fMinDelay, "delay", defaultMinDelay, "Minimum delay (ms) between requests to XRPL server")
	return cmd
}

func (cmd *BackfillCommand) Init(args []string) error {
	err := cmd.fs.Parse(args)
	if err != nil {
		return err
	}

	return cmd.Validate()
}

func (cmd *BackfillCommand) Validate() error {
	// Ledgers are backfilled in chronological order. Therefore, --from ledger
	// index must be less than --to ledger index.
	if cmd.fIndexFrom > cmd.fIndexTo {
		return fmt.Errorf("from ledger (%d) must be less than to ledger (%d)", cmd.fIndexFrom, cmd.fIndexTo)
	}
	return nil
}

func (cmd *BackfillCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd *BackfillCommand) Run() error {
	// Register command line signal handlers to gracefully shutdown cli
	signals.HandleAll()

	// Load validated config file
	config.EnvLoad(cmd.fConfigFile)

	// If websocket url is not provided in the cli, use the url from environment
	if cmd.fXrplServer == "" {
		cmd.fXrplServer = config.EnvXrplWebsocketFullHistoryURL()
	}

	// Initialize connections to services
	logger.New()
	connections.NewWriter()
	connections.NewXrplClientWithURL(cmd.fXrplServer)
	defer connections.CloseWriter()
	defer connections.CloseXrplClient()

	// Fetch ledger and queue transactions for indexing
	for ledgerIndex := cmd.fIndexFrom; ledgerIndex <= cmd.fIndexTo; ledgerIndex++ {
		startTime := time.Now().UnixNano() / int64(time.Millisecond)
		cmd.backfillLedger(ledgerIndex)
		reqDuration := time.Now().UnixNano()/int64(time.Millisecond) - startTime

		// Honor fair usage policy and wait before sending next request
		delayRequired := cmd.fMinDelay - reqDuration
		if delayRequired > 0 {
			time.Sleep(time.Duration(delayRequired) * time.Millisecond)
		}
	}
	return nil
}

// Worker functions
func (cmd *BackfillCommand) backfillLedger(ledgerIndex int) {
	log.Printf("[%s] Backfilling ledger: %d\n", cmd.fXrplServer, ledgerIndex)
	ledger := models.LedgerStream{
		Type:        models.LEDGER_STREAM_TYPE,
		LedgerIndex: uint32(ledgerIndex),
	}
	ledgerJSON, _ := json.Marshal(ledger)
	producers.ProduceLedger(connections.KafkaWriter, ledgerJSON)
	cmd.backfillTransactions(ledgerJSON)
}

func (cmd *BackfillCommand) backfillTransactions(ledgerJSON []byte) {
	producers.ProduceTransactions(connections.KafkaWriter, ledgerJSON)
}
