package models

type Meta struct {
	TransactionResult string   `json:"TransactionResult,omitempty"`
	DeliveredAmount   Currency `json:"DeliveredAmount,omitempty"`
	Delivered_Amount  Currency `json:"delivered_amount,omitempty"`
	TransactionIndex  uint32   `json:"TransactionIndex,omitempty"`
}
