package issuer

import (
	"encoding/json"
	"math/big"
	"time"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	ErrUnexpectedStatusCode = errors.New("unexpected status code")
)

type UUIDResponse struct {
	Id string `json:"id"`
}

type CredentialRequest struct {
	CredentialSchema  string            `json:"credentialSchema"`
	Type              string            `json:"type"`
	CredentialSubject CredentialSubject `json:"credentialSubject"`
	Expiration        *time.Time        `json:"expiration,omitempty"`
	MtProof           bool              `json:"mtProof"`
	SignatureProof    bool              `json:"signatureProof"`
}

type CredentialStatus struct {
	RevocationNonce int64 `json:"revocationNonce"`
}

type CredentialSubject struct {
	ID                string   `json:"id"`
	IsAdult           bool     `json:"isAdult"`
	IssuingAuthority  int64    `json:"issuingAuthority"`
	DocumentNullifier *big.Int `json:"documentNullifier"`
	CredentialHash    *big.Int `json:"credentialHash"`
	UserID            string   `json:"userid"`
	Features          string   `json:"f"`
	UserAddress       string   `json:"pk"`
	Metadata          string   `json:"metadata"`
}

type GetCredentialResponse struct {
	Id                    string           `json:"id"`
	ProofTypes            []string         `json:"proofTypes"`
	CreatedAt             time.Time        `json:"createdAt"`
	ExpiresAt             time.Time        `json:"expiresAt"`
	Expired               bool             `json:"expired"`
	SchemaHash            string           `json:"schemaHash"`
	SchemaType            string           `json:"schemaType"`
	SchemaUrl             string           `json:"schemaUrl"`
	Revoked               bool             `json:"revoked"`
	CredentialStatus      CredentialStatus `json:"credentialStatus"`
	CredentialSubject     json.RawMessage  `json:"credentialSubject"`
	UserID                string           `json:"userID"`
	SchemaTypeDescription string           `json:"schemaTypeDescription"`
}
