package user_test

import (
	"github.com/stretchr/testify/assert"
	"graphql-assent/pkg/user"
	"graphql-assent/pkg/user/mocks"
	"testing"
)

func TestModel_TableName(t *testing.T) {
	a := assert.New(t)

	m := user.Model{}
	tn := m.TableName()
	a.Equal("users", tn)
}

func TestModel_BeforeSave_NoPass(t *testing.T) {
	a := assert.New(t)

	m := user.Model{NewPassword: ""}

	err := m.BeforeSave()

	a.Nil(err)
	a.Zero(m.Password)
}

func TestModel_BeforeSave_Pass(t *testing.T) {
	a := assert.New(t)

	pa := new(mocks.PwdAlgo)
	pa.On("Hash", "test").Return("hash", nil)
	user.Services.Pwd = pa

	m := user.Model{NewPassword: "test"}

	err := m.BeforeSave()

	a.Nil(err)
	a.Equal("hash", m.Password)
	pa.AssertExpectations(t)
}
