package handler

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Handler interface {
	AddProduct(productName string, price float64, stock int) error
	UpdateProduct(productId int, productName string, price float64, stock int) error
	DeleteProduct(productName string) error
	CustomersTransactionsReport() error
}

type HandlerImpl struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *HandlerImpl {
	return &HandlerImpl{
		DB: db,
	}
}

func (h *HandlerImpl) AddProduct(productName string, price float64, stock int) error {
	_, err := h.DB.Exec("INSERT INTO Products (ProductName, Price, Stock) VALUES (?, ?, ?);", productName, price, stock)
	if err != nil {
		log.Print("Error inserting product to database: ", err)
		return err
	}

	log.Print("Successfully add new product")
	return nil
}

func (h *HandlerImpl) UpdateProduct(productId int, productName string, price float64, stock int) error {
	_, err := h.DB.Exec("UPDATE Products SET ProductName = ?, Price = ?, Stock = ? WHERE productId = ?;", productName, price, stock, productId)
	if err != nil {
		log.Print("Error updating product: ", err)
		return err
	}

	log.Print("Successfully update product")
	return nil
}

func (h *HandlerImpl) DeleteProduct(productName string) error {
	var exists bool
	err := h.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM Products WHERE name=?)", productName).Scan(&exists)
	if err != nil {
		log.Print("Error checking product existence: ", err)
		return err
	}

	if !exists {
		log.Print("Product does not exist")
		return fmt.Errorf("product with name '%s' does not exist", productName)
	} else {
		_, err = h.DB.Exec("DELETE FROM Products WHERE name=?", productName)
		if err != nil {
			log.Print("Error deleting record: ", err)
			return err
		}

		log.Print("Product deleted successfully")
	}

	return nil
}
