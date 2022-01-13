package accounts

import (
	"context"
	"fmt"

	"github.com/H15Z/-interview-accountapi/client/models"
	"github.com/google/uuid"
)

func (a *Accounts) Create(ctx context.Context, d models.AccountData) (string, error) {

	//generate guid for request
	guid := uuid.New().String()
	d.ID = guid

	b, err := a.Client.PostRequest(ctx, "/organisation/accounts", d)

	// handle response
	fmt.Println(string(b))

	return guid, err
}
