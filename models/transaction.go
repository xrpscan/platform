package models

import (
	"fmt"

	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/xrpl-go"
)

type Transaction struct {
	Hash               string `json:"hash,omitempty"`
	Account            string `json:"Account,omitempty"`
	TransactionType    string `json:"TransactionType,omitempty"`
	Fee                string `json:"Fee,omitempty"`
	Sequence           string `json:"Sequence,omitempty"`
	AccountTxnID       string `json:"AccountTxnID,omitempty"`
	Flags              int    `json:"Flags,omitempty"`
	LastLedgerSequence int    `json:"LastLedgerSequence,omitempty"`
	SourceTag          int    `json:"SourceTag,omitempty"`
	SigningPubKey      string `json:"SigningPubKey,omitempty"`
	TicketSequence     int    `json:"TicketSequence,omitempty"`
	TxnSignature       string `json:"TxnSignature,omitempty"`
	Date               int    `json:"date,omitempty"`
	LedgerIndex        int    `json:"ledger_index,omitempty"`
	Validated          bool   `json:"validated,omitempty"`
}

func FetchTransaction(ledgerIndex string) (xrpl.BaseResponse, error) {
	requestId := fmt.Sprintf("ledger.%s.tx", ledgerIndex)
	request := xrpl.BaseRequest{
		"id":           requestId,
		"command":      "ledger",
		"ledger_index": ledgerIndex,
		"transactions": true,
		"expand":       true,
	}
	response, err := connections.XrplClient.Request(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
