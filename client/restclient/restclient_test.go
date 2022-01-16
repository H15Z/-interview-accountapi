package restclient

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/H15Z/-interview-accountapi/client/models"
	"github.com/stretchr/testify/assert"
)

// test post
func TestRestPOST(t *testing.T) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c := NewClient(Defaults())

	// 404 not found
	resp, err := c.PostRequest(ctx, "", models.AccountData{})
	assert.Equal(t, "Page not found", err.Error())
	assert.Equal(t, `{"code":"PAGE_NOT_FOUND","message":"Page not found"}`, string(resp))

	// valid url
	_, err = c.PostRequest(ctx, "/v1/organisation/accounts", models.AccountData{})
	assert.Equal(t, "validation failure list:\nvalidation failure list:\nattributes in body is required\nid in body is required\norganisation_id in body is required\ntype in body is required", err.Error())

}

// test get
func TestRestGET(t *testing.T) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c := NewClient(Defaults())

	// 404 not found
	resp, err := c.GetRequest(ctx, "")
	assert.Equal(t, "Page not found", err.Error())
	assert.Equal(t, `{"code":"PAGE_NOT_FOUND","message":"Page not found"}`, string(resp))

	resp, err = c.GetRequest(ctx, "/v1/organisation/accounts")
	assert.Equal(t, nil, err)
	assert.True(t, strings.Contains(string(resp), `{"data":[{"attributes":{"`))

}

// test delete
func TestRestDELETE(t *testing.T) {

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	c := NewClient(Defaults())

	// 404 not found
	resp, err := c.DeleteRequest(ctx, "")
	assert.Equal(t, "delete failed with status code: 404 Not Found", err.Error())
	assert.Equal(t, resp, 404)

	// 400
	resp, err = c.DeleteRequest(ctx, "/v1/organisation/accounts/771d850a-f6b3-4c16-b544-bbc8a05b740d")
	assert.Equal(t, resp, 400)

	fmt.Println(resp, err)
}
