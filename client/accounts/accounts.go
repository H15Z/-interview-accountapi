package accounts

import (
	"encoding/json"

	"github.com/H15Z/-interview-accountapi/client/models"
	"github.com/H15Z/-interview-accountapi/client/ports"
	"github.com/google/uuid"
)

type Accounts struct {
	Client ports.RestClient
}

func NewAccounts(c ports.RestClient) *Accounts {
	return &Accounts{
		Client: c,
	}
}

// parse API response
func (a Accounts) parseResponse(b []byte) (models.AccountsResponse, error) {
	var r models.AccountsResponse
	err := json.Unmarshal(b, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
