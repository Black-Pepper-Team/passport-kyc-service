package requests

import (
	"encoding/json"
	snarkTypes "github.com/iden3/go-rapidsnark/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
)

type CreateIdentityRequestData struct {
	ID        string             `json:"id"`
	ZKProof   snarkTypes.ZKProof `json:"zkproof"`
	IDCardSOD struct {
		Asn1Data  string `json:"asn1data"`
		Algorithm string `json:"algorithm"`
		Signature string `json:"signature"`
		PemFile   string `json:"pemFile"`
	} `json:"id_card_sod"`
}

type CreateIdentityRequest struct {
	Data CreateIdentityRequestData `json:"data"`
}

func NewCreateIdentityRequest(r *http.Request) (CreateIdentityRequest, error) {
	var request CreateIdentityRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, nil
}