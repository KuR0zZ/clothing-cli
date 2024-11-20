package handler

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Handler interface {
	AddProduct(productName string, price float64, stock int) error
	ShowAllProducts() error
	UpdateProduct(productId int, productName string, price float64, stock int) error
	DeleteProduct(productName string) error
	CustomersTransactionsReport() error
	CurrentStockReport() error
	TotalRevenueReport() error
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

func (h *HandlerImpl) CustomersTransactionsReport() error {
	rows, err := h.DB.Query("SELECT Customers.Name, COUNT(Transactions.Id) AS NumberOfTransaction FROM Customers INNER JOIN Transactions ON Customers.Id = Transactions.CustomerId GROUP BY Customers.Id ORDER BY NumberOfTransaction DESC;")
	if err != nil {
		log.Print("Error fetching report: ", err)
		return err
	}
	defer rows.Close()

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

// create function to show all products
func (h *HandlerImpl) ShowAllProducts() error {
	rows, err := h.DB.Query("SELECT * FROM Products;")
	if err != nil {
		log.Print("Error fetching products: ", err)
		return err
	}

	fmt.Println("ID\tProduct Name\tPrice\tStock")
	fmt.Println("---------------------------------")

	defer rows.Close()

	for rows.Next() {
		var id int
		var productName string
		var price float64
		var stock int

		err = rows.Scan(&id, &productName, &price, &stock)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}

		fmt.Printf("%-8d %-24s %-12.2f %-6d\n", id, productName, price, stock)
	}

	return nil
}

// create function report for current stock
func (h *HandlerImpl) CurrentStockReport() error {
	rows, err := h.DB.Query("SELECT ProductName, Stock FROM Products ORDER BY Stock Asc;")
	if err != nil {
		log.Print("Error fetching report: ", err)
		return err
	}

	fmt.Println("Product Name\tStock")

	defer rows.Close()

	for rows.Next() {
		var productName string
		var stock int

		err = rows.Scan(&productName, &stock)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}

		fmt.Printf("%s\t%d\n", productName, stock)
	}

	return nil
}

func (h *HandlerImpl) TotalRevenueReport() error {
	rows, err := h.DB.Query(`
		SELECT p.ProductName, SUM(t.Price) AS TotalRevenue
		FROM Products p
		JOIN TransactionsDetails t ON p.id = t.id
		GROUP BY p.ProductName
		ORDER BY TotalRevenue DESC;
		`)
	if err != nil {
		log.Print("Error fetching records: ", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var productName string
		var totalRevenue float64

		err = rows.Scan(&productName, &totalRevenue)
		if err != nil {
			log.Print("Error scanning record: ", err)
			return err
		}

		fmt.Printf("%s\t%.2f\n", productName, totalRevenue)
	}

	return nil
}
