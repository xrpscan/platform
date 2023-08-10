package models

type SignerEntry struct {
	Account       string `json:"Account,omitempty"`
	SignerWeight  uint16 `json:"SignerWeight,omitempty"`
	WalletLocator string `json:"WalletLocator,omitempty"`
}
