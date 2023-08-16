package models

import (
	"errors"
	"fmt"

	"github.com/xrpscan/platform/connections"
	"github.com/xrpscan/xrpl-go"
	"github.com/xrpscan/xrpl-go/models"
)

// XRPL Genesis ledger is 32570 - https://xrpscan.com/ledger/32570
const GENESIS_LEDGER uint32 = 32570

// LedgerStream type is constant 'ledgerClosed' - https://xrpl.org/subscribe.html#ledger-stream
const LEDGER_STREAM_TYPE string = "ledgerClosed"

// Ledger struct represents output of 'ledger' websocket command
// Ref: https://xrpl.org/ledger.html#response-format
type Ledger struct {
	Accepted            bool                 `json:"accepted,omitempty"`
	AccountHash         string               `json:"account_hash,omitempty"`
	CloseFlags          uint32               `json:"close_flags,omitempty"`
	CloseTime           uint32               `json:"close_time,omitempty"`
	CloseTimeHuman      string               `json:"close_time_human,omitempty"`
	CloseTimeResolution uint32               `json:"close_time_resolution,omitempty"`
	Closed              bool                 `json:"closed,omitempty"`
	Hash                string               `json:"hash,omitempty"`
	LedgerHash          string               `json:"ledger_hash,omitempty"`
	LedgerIndex         uint32               `json:"ledger_index,omitempty"`
	ParentCloseTime     uint32               `json:"parent_close_time,omitempty"`
	ParentHash          string               `json:"parent_hash,omitempty"`
	SeqNum              uint32               `json:"seq_num,omitempty"`
	TotalCoins          string               `json:"totalCoins,omitempty"`
	Total_Coins         string               `json:"total_coins,omitempty"`
	TransactionHash     string               `json:"transaction_hash,omitempty"`
	Transactions        []models.Transaction `json:"transactions,omitempty"`
}

// LedgerStream struct represents ledger object emitted by ledger stream
// Ref: https://xrpl.org/subscribe.html#ledger-stream
type LedgerStream struct {
	Type             string `json:"type,omitempty"`
	FeeBase          uint64 `json:"fee_base,omitempty"`
	FeeRef           uint64 `json:"fee_ref,omitempty"`
	LedgerHash       string `json:"ledger_hash,omitempty"`
	LedgerIndex      uint32 `json:"ledger_index,omitempty"`
	LedgerTime       uint32 `json:"ledger_time,omitempty"`
	ReserveBase      uint64 `json:"reserve_base,omitempty"`
	ReserveInc       uint64 `json:"reserve_inc,omitempty"`
	TxnCount         uint32 `json:"txn_count,omitempty"`
	ValidatedLedgers string `json:"validated_ledgers,omitempty"`
}

func (ledger *LedgerStream) Validate() error {
	if ledger.Type != LEDGER_STREAM_TYPE {
		return errors.New("invalid LedgerStream object")
	}
	if ledger.LedgerIndex < GENESIS_LEDGER {
		return errors.New("invalid ledger_index")
	}
	if len(ledger.LedgerHash) != 64 {
		return errors.New("invalid ledger_hash")
	}
	return nil
}

// Fetches all transaction for a specific ledger from rippled
func (ledger *LedgerStream) FetchTransactions() (xrpl.BaseResponse, error) {
	if err := ledger.Validate(); err != nil {
		return nil, errors.New("invalid ledger_index")
	}

	requestId := fmt.Sprintf("ledger.%v.tx", ledger.LedgerIndex)
	request := xrpl.BaseRequest{
		"id":           requestId,
		"command":      "ledger",
		"ledger_index": ledger.LedgerIndex,
		"transactions": true,
		"expand":       true,
	}
	response, err := connections.XrplClient.Request(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
