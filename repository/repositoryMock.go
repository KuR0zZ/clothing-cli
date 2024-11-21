package repository

import (
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock.Mock
}

// AddProduct implements Repository.
func (r *RepoMock) AddProduct(productName string, price float64, stock int) error {
	args := r.Called(productName, price, stock)
	return args.Error(0)
}

// CurrentStockReport implements Repository.
func (r *RepoMock) CurrentStockReport() error {
	panic("unimplemented")
}

// CustomersTransactionsReport implements Repository.
func (r *RepoMock) CustomersTransactionsReport() error {
	panic("unimplemented")
}

// DeleteProduct implements Repository.
func (r *RepoMock) DeleteProduct(productName string) error {
	panic("unimplemented")
}

// ShowAllProducts implements Repository.
func (r *RepoMock) ShowAllProducts() error {
	panic("unimplemented")
}

// TotalRevenueReport implements Repository.
func (r *RepoMock) TotalRevenueReport() error {
	panic("unimplemented")
}

// UpdateProduct implements Repository.
func (r *RepoMock) UpdateProduct(productId int, productName string, price float64, stock int) error {
	panic("unimplemented")
}

func (r *RepoMock) UserLogin(email, password string) (string, error) {
	args := r.Called(email, password)
	return args.String(0), args.Error(1)
}
