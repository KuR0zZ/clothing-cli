package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

type Repository interface {
	UserLogin(email, password string) (string, error)
	AddProduct(productName string, price float64, stock int) error
	// ShowAllProducts() error
	// UpdateProduct(productId int, productName string, price float64, stock int) error
	DeleteProduct(productName string) error
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

func (h *RepoImpl) AddProduct(productName string, price float64, stock int) error {

	_, err := h.DB.Query("INSERT INTO Products (ProductName, Price, Stock) VALUES ($1, $2, $3);", productName, price, stock)

	if err != nil {
		log.Print("Error inserting product to database: ", err)
		return err
	}

	return nil
}

func (h *RepoImpl) DeleteProduct(productName string) error {
	result, err := h.DB.Exec("DELETE FROM Products WHERE ProductName=$1;", productName)
	if err != nil {
		return fmt.Errorf("error deleting product '%s': %v", productName, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("product with name '%s' does not exist", productName)
	}

	return nil
}

// var exists bool
// err := h.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM Products WHERE ProductName=$1);", productName).Scan(&exists)

// if err != nil {
// 	log.Print("Error checking product existence: ", err)
// 	return err
// }

// if !exists {
// 	log.Print("Product does not exist")
// 	return fmt.Errorf("product with name '%s' does not exist", productName)
// } else {
// 	_, err = h.DB.Exec("DELETE FROM Products WHERE ProductName=$1;", productName)
// 	if err != nil {
// 		log.Print("Error deleting record: ", err)
// 		return err
// 	}
// }

// return nil
