package producers

import (
	"fmt"

	"github.com/xrpscan/platform/connections"
)

func ReadLedgerStream() {
	fmt.Println("Reading validated ledgers")
	connections.XrplClient.Subscribe("ledger")
	connections.XrplClient.Subscribe("transactions")
}
