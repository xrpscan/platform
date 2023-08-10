package models

type Signer struct {
	Account       string `json:"Account,omitempty"`
	TxnSignature  string `json:"TxnSignature,omitempty"`
	SigningPubKey string `json:"SigningPubKey,omitempty"`
}
