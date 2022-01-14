package client

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/H15Z/-interview-accountapi/client/models"
	"github.com/stretchr/testify/assert"
)

var testing_accounts_resource_link string
var testing_accounts_version int

func TestCreateClient(t *testing.T) {
	c := New()
	assert.NotEmpty(t, c.Accounts.Client)

	//TODO add more assetions here
}

//End-to-end accounts tests
//RUN Tests in sequence
func TestAccountsCreateSuccess(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c := New()

	/*
		REFERENCE DATA

		{
		"data": {
			"id": "771d850a-f6b3-4c16-b544-bbc8a05b740d",
			"organisation_id": "e2e47bf9-1d1a-46fc-930d-92a43903f857",
			"type": "accounts",
			"attributes": {
				"country": "GB",
				"base_currency": "GBP",
				"bank_id": "400302",
				"bank_id_code": "GBDSC",
				"account_number": "10000004",
				"customer_id": "234",
				"iban": "GB28NWBK40030212764204",
				"bic": "NWBKGB42",
				"account_classification": "Personal",
				"name" : ["name"]
			}
		}

	*/

	//Declare Payload Struct
	country := "GB"
	account_classification := "Personal"

	payload := models.AccountData{
		OrganisationID: "e2e47bf9-1d1a-46fc-930d-92a43903f857",
		Type:           "accounts",
		Attributes: &models.AccountAttributes{
			Country:               &country,
			BaseCurrency:          "GBP",
			BankID:                "400302",
			BankIDCode:            "GBDSC",
			AccountNumber:         "10000004",
			Iban:                  "GB28NWBK40030212764204",
			Bic:                   "NWBKGB42",
			AccountClassification: &account_classification,
			Name:                  []string{"Name", "Surname"},
			// CustomerId:    "234",
		},
	}

	guid, resp, err := c.Accounts.Create(ctx, payload)
	assert.Equal(t, nil, err)
	// assert.Equal(t, "/v1/organisation/accounts/"+guid, resource_link)
	assert.True(t, strings.Contains(resp.Links.Self, "/v1/organisation/accounts/"))
	assert.NotEmpty(t, resp.Data)
	assert.NotEqual(t, "", guid)
	// assert.True(t, IsValidUUID(guid))

	testing_accounts_resource_link = resp.Links.Self
	testing_accounts_version = int(*resp.Data.Version)

}

func TestAccountsCreateFailWithError(t *testing.T) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c := New()

	//Declare Payload Struct
	country := "GB"
	account_classification := "Personal"

	payload := models.AccountData{
		OrganisationID: "",
		Type:           "accounts",
		Attributes: &models.AccountAttributes{
			Country:               &country,
			BaseCurrency:          "GBP",
			BankID:                "400302",
			BankIDCode:            "GBDSC",
			AccountNumber:         "10000004",
			Bic:                   "NWBKGB42",
			AccountClassification: &account_classification,
			// Iban:                  "GB28NWBK40030212764204",
			// Name:                  []string{"Name", "Surname"},
			// CustomerId:    "234",
		},
	}

	_, _, err := c.Accounts.Create(ctx, payload)
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "validation failure list:\nvalidation failure list:\nvalidation failure list:\nname in body is required\norganisation_id in body is required", err.Error())

}

func TestAccountsFetchSuccess(t *testing.T) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c := New()

	resp, err := c.Accounts.FetchByResource(ctx, testing_accounts_resource_link)

	assert.Equal(t, nil, err)
	assert.Equal(t, testing_accounts_resource_link, resp.Links.Self)
	assert.Equal(t, "GB28NWBK40030212764204", resp.Data.Attributes.Iban)
	assert.NotEmpty(t, resp.Data)

}

func TestAccountsFetchFail(t *testing.T) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c := New()

	resp, err := c.Accounts.FetchByResource(ctx, "/v1/organisation/accounts/771d850a-f6b3-4c16-b544-bbc8a05b740d")

	assert.NotEqual(t, nil, err)
	assert.Equal(t, "record 771d850a-f6b3-4c16-b544-bbc8a05b740d does not exist", err.Error())
	assert.Empty(t, resp.Data)

}

// delete fail
func TestAccountsDeleteFailWrongVersion(t *testing.T) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c := New()

	_, err := c.Accounts.Delete(ctx, testing_accounts_resource_link, 99)
	assert.Equal(t, "delete failed with status code: 409 Conflict", err.Error())
}

// delete success
func TestAccountsDeleteSuccess(t *testing.T) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c := New()

	resp, err := c.Accounts.Delete(ctx, testing_accounts_resource_link, testing_accounts_version)

	assert.Equal(t, nil, err)
	assert.Equal(t, "204\tNo Content\tResource has been successfully deleted", resp)

}
