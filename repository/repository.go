package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Repository interface {
	UserLogin(email, password string) (string, error)
	// AddProduct(productName string, price float64, stock int) error
	// ShowAllProducts() error
	// UpdateProduct(productId int, productName string, price float64, stock int) error
	// DeleteProduct(productName string) error
	// CustomersTransactionsReport() error
	// CurrentStockReport() error
	// TotalRevenueReport() error
}

type RepoImpl struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *RepoImpl {
	return &RepoImpl{
		DB: db,
	}
}

func (h *RepoImpl) UserLogin(email, password string) (string, error) {
	var dbPassword string

	err := h.DB.QueryRow("SELECT Password FROM UserAdmin WHERE Email=$1;", email).Scan(&dbPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("user with email %s not found", email)
		}
		return "", fmt.Errorf("error querying user admin: %v", err)
	}

	return dbPassword, nil

}
