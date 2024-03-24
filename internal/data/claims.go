package data

import (
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
)

type ClaimQ interface {
	New() ClaimQ
	Insert(value Claim) error
	FilterBy(column string, value any) ClaimQ
	Get() (*Claim, error)
	Select() ([]Claim, error)
	DeleteByID(id uuid.UUID) error
	ForUpdate() ClaimQ
	ResetFilter() ClaimQ
}

type Claim struct {
	ID           uuid.UUID      `db:"id"            structs:"id"`
	UserID       uuid.UUID      `db:"user_id"       structs:"user_id"`
	UserDID      string         `db:"user_did"      structs:"user_did"`
	IssuerDID    string         `db:"issuer_did"    structs:"issuer_did"`
	UserAddress  common.Address `db:"user_address"  structs:"user_address"`
	DocumentHash string         `db:"document_hash" structs:"document_hash"`
	CreatedAt    time.Time      `db:"created_at"    structs:"-"`
}
