package accounts

import (
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/H15Z/-interview-accountapi/client/models"
	"github.com/stretchr/testify/assert"
)

// test accounts by injecting stub structs
// decoupled from rest client, allows for independent testing

func TestCreateSuccess(t *testing.T) {
	// create stub rest client
	stub_client := stubClient{
		PostResponse: stubResponse{
			Error: nil,
			Body:  testing_stub_response, //mock of succesfull response
		},
	}

	accounts := NewAccounts(stub_client)
	guid, resp, err := accounts.Create(context.Background(), models.AccountData{})

	assert.True(t, strings.Contains(resp.Links.Self, "/v1/organisation/accounts/"))
	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", guid)
}

func TestCreateFailure(t *testing.T) {
	// create stub rest client
	stub_client := stubClient{
		PostResponse: stubResponse{
			Error: errors.New("Mock testing error"),
			Body:  []byte(`{"error_message" : "Mock testing error"}`),
		},
	}

	accounts := NewAccounts(stub_client)
	guid, resp, err := accounts.Create(context.Background(), models.AccountData{})

	assert.Equal(t, "Mock testing error", err.Error())
	assert.NotEqual(t, "", guid)
	assert.Equal(t, "", resp.Links.Self)

}

func TestFetchSucces(t *testing.T) {
	// create stub rest client
	stub_client := stubClient{
		GetResponse: stubResponse{
			Error: nil,
			Body:  testing_stub_response,
		},
	}

	accounts := NewAccounts(stub_client)
	resource_link := "/v1/organisation/accounts/771d850a-f6b3-4c16-b544-bbc8a05b740d"
	resp, err := accounts.FetchByResource(context.Background(), resource_link)

	assert.Equal(t, nil, err)
	assert.Equal(t, resource_link, resp.Links.Self)
	assert.NotEmpty(t, resp.Data)
	assert.Equal(t, "e2e47bf9-1d1a-46fc-930d-92a43903f857", resp.Data.OrganisationID)

}

//fetch by id

//fetch failure
func TestFetchFail(t *testing.T) {
	// create stub rest client
	stub_client := stubClient{
		GetResponse: stubResponse{
			Error: errors.New("record 771d850a-f6b3-4c16-b544-bbc8a05b740d does not exist"),
			Body:  []byte(`{}`),
		},
	}

	accounts := NewAccounts(stub_client)
	resource_link := "/v1/organisation/accounts/771d850a-f6b3-4c16-b544-bbc8a05b740d"
	resp, err := accounts.FetchByResource(context.Background(), resource_link)

	assert.Equal(t, "record 771d850a-f6b3-4c16-b544-bbc8a05b740d does not exist", err.Error())
	assert.Empty(t, resp.Data)

}

//delete success
func TestDeleteSuccess(t *testing.T) {
	// create stub rest client
	stub_client := stubClient{
		DeleteResponse: stubResponse{
			ResponseCode: 204,
			Error:        nil,
		},
	}

	accounts := NewAccounts(stub_client)
	resource_link := "/v1/organisation/accounts/771d850a-f6b3-4c16-b544-bbc8a05b740d"
	resp, err := accounts.Delete(context.Background(), resource_link, 1)

	assert.Equal(t, nil, err)
	assert.Equal(t, "204\tNo Content\tResource has been successfully deleted", resp)

}

//delete failures
func TestDeleteFail404(t *testing.T) {
	// create stub rest client
	stub_client := stubClient{
		DeleteResponse: stubResponse{
			ResponseCode: 404,
			Error:        errors.New("delete fail error"),
		},
	}

	accounts := NewAccounts(stub_client)
	resource_link := "/v1/organisation/accounts/771d850a-f6b3-4c16-b544-bbc8a05b740d"
	resp, err := accounts.Delete(context.Background(), resource_link, 1)

	assert.Equal(t, "delete fail error", err.Error())
	assert.Equal(t, "404\tNot Found\tSpecified resource does not exist", resp)

}
func TestDeleteFail409(t *testing.T) {
	// create stub rest client
	stub_client := stubClient{
		DeleteResponse: stubResponse{
			ResponseCode: 409,
			Error:        errors.New("delete fail error"),
		},
	}

	accounts := NewAccounts(stub_client)
	resource_link := "/v1/organisation/accounts/771d850a-f6b3-4c16-b544-bbc8a05b740d"
	resp, err := accounts.Delete(context.Background(), resource_link, 1)

	assert.Equal(t, "delete fail error", err.Error())
	assert.Equal(t, "409\tConflict\tSpecified version incorrect", resp)

}

var testing_stub_response []byte = []byte(`
	{
		"data":
		{
			"attributes":
			{
				"account_classification": "Personal",
				"account_number": "10000004",
				"alternative_names": null,
				"bank_id": "400302",
				"bank_id_code": "GBDSC",
				"base_currency": "GBP",
				"bic": "NWBKGB42",
				"country": "GB",
				"iban": "GB28NWBK40030212764204",
				"name":
				[
					"name"
				]
			},
			"created_on": "2022-01-13T15:37:36.053Z",
			"id": "771d850a-f6b3-4c16-b544-bbc8a05b740d",
			"modified_on": "2022-01-13T15:37:36.053Z",
			"organisation_id": "e2e47bf9-1d1a-46fc-930d-92a43903f857",
			"type": "accounts",
			"version": 0
		},
		"links":
		{
			"self": "/v1/organisation/accounts/771d850a-f6b3-4c16-b544-bbc8a05b740d"
		}
	}`)
