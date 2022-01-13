package accounts

import (
	"context"
	"errors"
	"strings"
)

func (a *Accounts) Delete(ctx context.Context, resource string) error {
	if !strings.Contains(resource, "/v1/organisation/accounts/") {
		return errors.New("account API resource link not valid")
	}

	return a.Client.DeleteRequest(ctx, resource)
}
