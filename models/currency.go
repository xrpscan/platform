package models

type Currency struct {
	Currency string `json:"currency,omitempty"`
	Issuer   string `json:"issuer,omitempty"`
	Value    int64  `json:"value,omitempty"`
}
