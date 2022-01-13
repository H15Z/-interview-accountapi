package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatClient(t *testing.T) {
	c := New()

	assert.Equal(t, c.Accounts.Client, "")
}
