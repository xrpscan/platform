package models

type Meta struct {
	// AffectedNodes // Unimplemented
	DeliveredAmount   Currency `json:"DeliveredAmount,omitempty"`
	Delivered_Amount  Currency `json:"delivered_amount,omitempty"`
	TransactionIndex  int      `json:"TransactionIndex,omitempty"`
	TransactionResult string   `json:"TransactionResult,omitempty"`
}
