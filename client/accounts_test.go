package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountsCreate(t *testing.T) {
	c := New()

	err := c.Accounts.Create()

	assert.Equal(t, err, nil)
}
