package models

const XRP = "XRP"
const DROPS_IN_XRP = 1000000

type Currency struct {
	Currency string `json:"currency,omitempty"`
	Issuer   string `json:"issuer,omitempty"`
	Value    int64  `json:"value,omitempty"`
	Native   bool   `json:"native,omitempty"`
}
