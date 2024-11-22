package handler

import (
	"clothing-cli/repository"
	"errors"
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
type AddProductTest struct {
	name  string
	price float64
	stock int
}
type DeleteProductTest struct {
	name string
}

func TestUserLogin(t *testing.T) {
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

func TestAddProduct(t *testing.T) {
	// Arrange
	product := AddProductTest{
		name:  "TestProduct",
		price: 99999,
		stock: 10,
	}

	// Mock the repository method
	repoMock.On("AddProduct", product.name, product.price, product.stock).Return(nil)

	// Act
	err := userHandler.AddProduct(product.name, product.price, product.stock)

	// Assert
	assert.Nil(t, err)                                                                 // Ensure no error occurred
	repoMock.AssertCalled(t, "AddProduct", product.name, product.price, product.stock) // Ensure method was called
}

func TestAddProductFail(t *testing.T) {
	// Arrange
	product := AddProductTest{
		name:  "TestProduct",
		price: 99999,
		stock: 10,
	}

	expectedErr := errors.New("database error")

	// Mock the repository method to return an error
	repoMock.On("AddProduct", product.name, product.price, product.stock).Return(expectedErr)

	// Act
	err := userHandler.AddProduct(product.name, product.price, product.stock)

	// Assert
	assert.NotNil(t, err)                                                              // Ensure an error occurred
	assert.Equal(t, expectedErr, err)                                                  // Ensure the returned error matches the mock
	repoMock.AssertCalled(t, "AddProduct", product.name, product.price, product.stock) // Ensure method was called
}

func TestDeleteProduct(t *testing.T) {
	// Arrange
	product := DeleteProductTest{
		name: "TestProduct",
	}

	// Mock the repository method
	repoMock.On("DeleteProduct", product.name).Return(nil)

	// Act
	err := userHandler.DeleteProduct(product.name)

	// Assert
	assert.Nil(t, err)                                      // Ensure no error occurred
	repoMock.AssertCalled(t, "DeleteProduct", product.name) // Ensure method was called
}

func TestDeleteProductFail(t *testing.T) {
	// Arrange
	product := DeleteProductTest{
		name: "TestProduct",
	}

	expectedErr := errors.New("database error")

	// Mock the repository method to return an error
	repoMock.On("DeleteProduct", product.name).Return(expectedErr)

	// Act
	err := userHandler.DeleteProduct(product.name)

	// Assert
	assert.NotNil(t, err)                                   // Ensure an error occurred
	assert.Equal(t, expectedErr, err)                       // Ensure the returned error matches the mock
	repoMock.AssertCalled(t, "DeleteProduct", product.name) // Ensure method was called
}
