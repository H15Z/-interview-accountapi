package accounts

import "github.com/H15Z/-interview-accountapi/client/ports"

type Accounts struct {
	Client ports.RestClient
}

func NewAccounts(c ports.RestClient) *Accounts {
	return &Accounts{
		Client: c,
	}
}
