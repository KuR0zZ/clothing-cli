package handler

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Handler interface {
	AddProduct(productName string, price float64, stock int) error
	UpdateProduct(productId int, productName string, price float64, stock int) error
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
