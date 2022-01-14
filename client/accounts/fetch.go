package accounts

import (
	"context"
	"errors"
	"strings"

	"github.com/H15Z/-interview-accountapi/client/models"
)

//Create account using Form3 API: https://api-docs.form3.tech/api.html#organisation-accounts-fetch
//Accepts context and account guid, resource link is generated from guid
//Returns accountsData struct and error
func (a Accounts) FetchById(ctx context.Context, guid string) (models.AccountsResponse, error) {
	//validate guid
	if !IsValidUUID(guid) {
		return models.AccountsResponse{}, errors.New("account GUID not valid")
	}

	resource := "/v1/organisation/accounts/" + guid
	return a.doFetch(ctx, resource)
}

//Create account using Form3 API: https://api-docs.form3.tech/api.html#organisation-accounts-fetch
//Accepts context and account resource link
//Returns accountsData struct and error
func (a Accounts) FetchByResource(ctx context.Context, resource string) (models.AccountsResponse, error) {
	//validate resource link
	if !strings.Contains(resource, "/v1/organisation/accounts/") {
		return models.AccountsResponse{}, errors.New("account API resource link not valid")
	}

	return a.doFetch(ctx, resource)
}

//common fetch function
func (a Accounts) doFetch(ctx context.Context, resource string) (models.AccountsResponse, error) {

	b, err := a.Client.GetRequest(ctx, resource)
	if err != nil {

		return models.AccountsResponse{}, err
	}
	return a.parseResponse(b)
}
