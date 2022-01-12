package client

import (
	"github.com/H15Z/-interview-accountapi/client/accounts"
	"github.com/H15Z/-interview-accountapi/client/restclient"
)

// client service
type client struct {
	Accounts *accounts.Accounts

	// other API resources can go here... or a separate service could be used
}

func New() *client {

	// create rest client and use defaults
	c := restclient.NewClient(restclient.Defaults())

	return &client{
		Accounts: accounts.NewAccounts(c),
	}
}
