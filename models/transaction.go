package models

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
