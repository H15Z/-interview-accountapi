package accounts

import (
	"context"
	"errors"
	"strconv"
	"strings"
)

//Deletes account version using Form3 API: https://api-docs.form3.tech/api.html#organisation-accounts-delete
//Accepts context, resource link for account and version number
//Returns error only as there is no response body to parse
func (a Accounts) Delete(ctx context.Context, resource string, version int) (string, error) {
	if !strings.Contains(resource, "/v1/organisation/accounts/") {
		return "Delete request failed", errors.New("account API resource link not valid")
	}

	code, err := a.Client.DeleteRequest(ctx, resource+"?version="+strconv.Itoa(int(version)))
	return a.handleDeleteResponse(code), err
}

func (a Accounts) handleDeleteResponse(code int) string {
	switch code {
	case 204:
		return "204	No Content	Resource has been successfully deleted"
	case 404:
		return "404	Not Found	Specified resource does not exist"
	case 409:
		return "409	Conflict	Specified version incorrect"
	default:
		return "Delete request failed"
	}
}
