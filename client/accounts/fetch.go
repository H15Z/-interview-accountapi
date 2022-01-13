package accounts

import (
	"context"
	"errors"
	"strings"

	"github.com/H15Z/-interview-accountapi/client/models"
)

func (a Accounts) Fetch(ctx context.Context, guid string) (models.AccountsResponse, error) {
	//validate guid
	if !IsValidUUID(guid) {
		return models.AccountsResponse{}, errors.New("account GUID not valid")
	}

	resource := "/v1/organisation/accounts/" + guid
	return a.doFetch(ctx, resource)
}

func (a Accounts) FetchByResource(ctx context.Context, resource string) (models.AccountsResponse, error) {
	//validate resource link
	if !strings.Contains(resource, "/v1/organisation/accounts/") {
		return models.AccountsResponse{}, errors.New("account API resource link not valid")
	}

	return a.doFetch(ctx, resource)
}

func (a Accounts) doFetch(ctx context.Context, resource string) (models.AccountsResponse, error) {
	var d models.AccountsResponse

	b, err := a.Client.GetRequest(ctx, resource)
	if err != nil {
		return d, err
	}

	return a.parseResponse(b)
}
