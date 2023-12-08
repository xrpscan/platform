package models

type Memos struct {
	Memo Memo `json:"Memo,omitempty"`
}

type Memo struct {
	MemoData   string `json:"MemoData,omitempty"`
	MemoFormat string `json:"MemoFormat,omitempty"`
	MemoType   string `json:"MemoType,omitempty"`
}
