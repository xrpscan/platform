package models

type AuthAccounts struct {
	AuthAccount AuthAccount `json:"AuthAccount,omitempty"`
}

type AuthAccount struct {
	Account string `json:"Account,omitempty"`
}
