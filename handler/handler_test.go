package handler

import (
	"clothing-cli/repository"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoMock = &repository.RepoMock{Mock: mock.Mock{}}
var userHandler = NewHandler(repoMock)

type UserLoginTest struct {
	email    string
	password string
}

func testUserLogin(t *testing.T) {
	user := UserLoginTest{
		email:    "santoso@gmail.com",
		password: "santoso123",
	}

	repoMock.On("UserLogin", user.email, user.password).Return(user.password, nil)

	result := userHandler.UserLogin(user.email, user.password)
	assert.Nil(t, result)
	assert.Equal(t, user.password, result)
}

func TestUserLoginFail(t *testing.T) {
	user := UserLoginTest{
		email:    "santoso@gmail.com",
		password: "santoso123",
	}

	repoMock.On("UserLogin", user.email, user.password).Return("", nil)

	result := userHandler.UserLogin(user.email, user.password)
	assert.NotNil(t, result)

	expectedError := fmt.Sprintf("incorrect password for email %s", user.email)

	assert.Equal(t, expectedError, result.Error())
}
