package models

const XRP = "XRP"

type Currency struct {
	Currency string `json:"currency,omitempty"`
	Issuer   string `json:"issuer,omitempty"`
	Value    int64  `json:"value,omitempty"`
	Native   bool   `json:"native,omitempty"`
}
