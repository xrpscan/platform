package types

type BaseRequest struct {
	Id         int    `json:"id,omitempty"`
	Command    string `json:"command,omitempty"`
	ApiVersion int16  `json:"api_version,omitempty"`
}

type Warning struct {
	Id      int               `json:"id,omitempty"`
	Message string            `json:"message,omitempty"`
	Details map[string]string `json:"details,omitempty"`
}

type BaseResponse struct {
	Id         int       `json:"id,omitempty"`
	Status     string    `json:"status,omitempty"`
	Type       string    `json:"type,omitempty"`
	Result     string    `json:"result,omitempty"`
	Warnings   []Warning `json:"warning,omitempty"`
	Forwarded  bool      `json:"forwarded,omitempty"`
	ApiVersion int16     `json:"api_version,omitempty"`
}

type ErrorResponse struct {
	Id           int    `json:"id,omitempty"`
	Status       string `json:"status,omitempty"`
	Type         string `json:"type,omitempty"`
	Error        string `json:"error,omitempty"`
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	ApiVersion   int16  `json:"api_version,omitempty"`
}
