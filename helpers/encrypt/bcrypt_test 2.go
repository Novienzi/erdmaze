package encrypt_test

import (
	"erdmaze/helpers/encrypt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "12345"
	hash, err := encrypt.Hash(password)

	assert.Nil(t, err)

	assert.True(t, encrypt.ValidateHash(password, hash))
}
