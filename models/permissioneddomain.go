package models

type AcceptedCredentials struct {
	Credential Credential `json:"Credential,omitempty"`
}

type Credential struct {
	Issuer         string `json:"Issuer,omitempty"`
	CredentialType string `json:"CredentialType,omitempty"`
}
