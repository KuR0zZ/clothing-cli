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

func (h *HandlerImpl) CustomersTransactionsReport() error {
	rows, err := h.DB.Query("SELECT Customers.Name, COUNT(Transactions.Id) AS NumberOfTransaction FROM Customers INNER JOIN Transactions ON Customers.Id = Transactions.CustomerId GROUP BY Customers.Id ORDER BY NumberOfTransaction DESC;")
	if err != nil {
		log.Print("Error fetching report: ", err)
		return err
	}

	fmt.Println("Name\tNumber Of Transaction")
	for rows.Next() {
		var customer_name string
		var number_of_transaction int

		err = rows.Scan(&customer_name, &number_of_transaction)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}

		fmt.Printf("%s\t%d", customer_name, number_of_transaction)
	}
	return nil
}
