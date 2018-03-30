package user_test

import (
	"github.com/stretchr/testify/assert"
	"graphql-assent/pkg/user"
	"graphql-assent/pkg/user/mocks"
	"testing"
)

func TestNewPwdAlgo(t *testing.T) {

	pa, err := user.NewPwdAlgo("argon2")

	assert.Nil(t, err)
	assert.IsType(t, user.PwdArgon2{}, pa)

	pa, err = user.NewPwdAlgo("e")

	assert.Nil(t, pa)
	assert.Error(t, err)
}

func TestPwdArgon2_Hash_Err(t *testing.T) {
	a2 := new(mocks.IArgon2)

	a2.On("GenerateSalt", 16).Return("", assert.AnError)

	pa := user.PwdArgon2{a2}

	h, err := pa.Hash("p")

	assert.Zero(t, h)
	assert.Error(t, err)
}

func TestPwdArgon2_Hash_OK(t *testing.T) {
	a := assert.New(t)
	a2 := new(mocks.IArgon2)

	a2.On("GenerateSalt", 16).Return("salt", nil)
	a2.On(
		"Key",
		[]byte("p"),
		[]byte("salt"),
		uint32(4),
		uint32(16384),
		uint8(4),
		uint32(32)).Return([]byte("hash"))

	a2.On("EncodeToString", []byte("hash")).Return("hash")

	pa := user.PwdArgon2{Crypto: a2}

	h, err := pa.Hash("p")

	a.Nil(err)
	a.Equal("Argon2$4$16384$4$32$salt$hash", h)
}
