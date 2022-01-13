package accounts

import (
	"context"

	"github.com/H15Z/-interview-accountapi/client/models"
	"github.com/google/uuid"
)

//Create account using Form3 API: https://api-docs.form3.tech/api.html#organisation-accounts-create
//Accepts context and  AccountData struct
//Returns guid of request , new resource link including guid of created account and any errors encountered
func (a Accounts) Create(ctx context.Context, d models.AccountData) (string, models.AccountsResponse, error) {

	//generate guid for request
	guid := uuid.New().String()
	d.ID = guid // assign guid to payload

	b, err := a.Client.PostRequest(ctx, "/organisation/accounts", d)
	if err != nil {
		return guid, models.AccountsResponse{}, err
	}

	// handle response
	r, err := a.parseResponse(b)
	if err != nil {
		return guid, r, err
	}

	return guid, r, err
}
