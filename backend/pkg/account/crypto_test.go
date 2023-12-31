package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_encryptPassword(t *testing.T) {
	t.Parallel()

	hash, err := encryptPassword("password")
	assert.NoError(t, err, "encryptPassword should not return an error")
	assert.NotEmpty(t, hash, "encryptPassword should return a hash")

	hash2, err := encryptPassword("password")
	assert.NoError(t, err, "encryptPassword should not return an error")
	assert.NotEmpty(t, hash2, "encryptPassword should return a hash")

	assert.NotEqual(t, hash, hash2, "encryptPassword should return different hashes for different passwords")
}

func Test_comparePassword(t *testing.T) {
	t.Parallel()

	hash, err := encryptPassword("password")
	assert.NoError(t, err, "encryptPassword should not return an error")
	assert.NotEmpty(t, hash, "encryptPassword should return a hash")

	err = comparePassword(hash, "password")
	assert.NoError(t, err, "comparePassword should not return an error")

	err = comparePassword(hash, "password2")
	assert.Error(t, err, "comparePassword should return an error")
}
