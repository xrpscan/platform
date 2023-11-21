package models

type SignerEntry struct {
	Account       string `json:"Account,omitempty"`
	WalletLocator string `json:"WalletLocator,omitempty"`
	SignerWeight  uint16 `json:"SignerWeight,omitempty"`
}
