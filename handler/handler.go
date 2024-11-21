package handler

import (
	"clothing-cli/repository"
	"fmt"
	"log"
)

type Handler struct {
	Repo repository.Repository
}

func NewHandler(repo repository.Repository) Handler {
	return Handler{repo}
}

func (h *Handler) UserLogin(email, password string) error {

	dbPassword, err := h.Repo.UserLogin(email, password)
	if err != nil {
		return err
	}

	if dbPassword != password {
		return fmt.Errorf("incorrect password for email %s", email)
	}

	log.Println("Login successful for user:", email)
	return nil
}

func (h *Handler) AddProduct(productName string, price float64, stock int) error {
	err := h.Repo.AddProduct(productName, price, stock)
	if err != nil {
		log.Print("Error inserting product to database: ", err)
		return err
	}

	log.Print("Successfully add new product")
	return nil
}

// func (h *HandlerImpl) UpdateProduct(productId int, productName string, price float64, stock int) error {
// 	var updatedId int
// 	err := h.DB.QueryRow("UPDATE Products SET ProductName = COALESCE(NULLIF($1, ''), ProductName), Price = COALESCE(NULLIF($2, 0), Price), Stock = COALESCE(NULLIF($3, 0), Stock) WHERE Id = $4 RETURNING Id;", productName, price, stock, productId).Scan(&updatedId)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return fmt.Errorf("product with id %d not found", productId)
// 		}
// 		return fmt.Errorf("error updating product: %v", err)
// 	}

// 	log.Printf("Successfully updated product with ID %d", updatedId)
// 	return nil
// }

// func (h *HandlerImpl) DeleteProduct(productName string) error {
// 	var exists bool
// 	err := h.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM Products WHERE ProductName=$1);", productName).Scan(&exists)
// 	if err != nil {
// 		log.Print("Error checking product existence: ", err)
// 		return err
// 	}

// 	if !exists {
// 		log.Print("Product does not exist")
// 		return fmt.Errorf("product with name '%s' does not exist", productName)
// 	} else {
// 		_, err = h.DB.Exec("DELETE FROM Products WHERE ProductName=$1;", productName)
// 		if err != nil {
// 			log.Print("Error deleting record: ", err)
// 			return err
// 		}

// 		log.Print("Product deleted successfully")
// 	}

// 	return nil
// }

// func (h *HandlerImpl) CustomersTransactionsReport() error {
// 	rows, err := h.DB.Query("SELECT Customers.Name, COUNT(Transactions.Id) AS NumberOfTransaction FROM Customers INNER JOIN Transactions ON Customers.Id = Transactions.CustomerId GROUP BY Customers.Id ORDER BY NumberOfTransaction DESC;")
// 	if err != nil {
// 		log.Print("Error fetching report: ", err)
// 		return err
// 	}
// 	defer rows.Close()

// 	fmt.Println("Name\t\tNumber Of Transaction")
// 	for rows.Next() {
// 		var customer_name string
// 		var number_of_transaction int

// 		err = rows.Scan(&customer_name, &number_of_transaction)
// 		if err != nil {
// 			log.Print("Error scanning record: ", err)
// 			return err
// 		}

// 		fmt.Printf("%s\t\t%d\n", customer_name, number_of_transaction)
// 	}
// 	return nil
// }

// // create function to show all products
// func (h *HandlerImpl) ShowAllProducts() error {
// 	rows, err := h.DB.Query("SELECT * FROM Products ORDER BY Id Asc;")
// 	if err != nil {
// 		log.Print("Error fetching products: ", err)
// 		return err
// 	}

// 	id := "ID"
// 	productName := "Product Name"
// 	price := "Price"
// 	stock := "Stock"

// 	fmt.Printf("%-3s %-24s %-12s %-6s\n", id, productName, price, stock)
// 	fmt.Println("--------------------------------------------------")

// 	defer rows.Close()

// 	for rows.Next() {
// 		var id int
// 		var productName string
// 		var price float64
// 		var stock int

// 		err = rows.Scan(&id, &productName, &price, &stock)
// 		if err != nil {
// 			log.Print("Error scanning record: ", err)
// 			return err
// 		}

// 		fmt.Printf("%-3d %-24s %-12.2f %-6d\n", id, productName, price, stock)
// 	}

// 	return nil
// }

// // create function report for current stock
// func (h *HandlerImpl) CurrentStockReport() error {
// 	rows, err := h.DB.Query("SELECT ProductName, Stock FROM Products ORDER BY Stock Asc;")
// 	if err != nil {
// 		log.Print("Error fetching report: ", err)
// 		return err
// 	}

// 	prodName := "Product Name"
// 	stock := "Stock"

// 	fmt.Printf("%-24s %-6s\n", prodName, stock)
// 	fmt.Println("--------------------------------")

// 	defer rows.Close()

// 	for rows.Next() {
// 		var productName string
// 		var stock int

// 		err = rows.Scan(&productName, &stock)
// 		if err != nil {
// 			log.Print("Error scanning record: ", err)
// 			return err
// 		}

// 		fmt.Printf("%-24s %-6d\n", productName, stock)
// 	}

// 	return nil
// }

// func (h *HandlerImpl) TotalRevenueReport() error {
// 	rows, err := h.DB.Query(`
// 		select
// 			p.ProductName,
// 			sum(td.TotalPrice) as TotalRevenue
// 		from
// 			TransactionsDetails td
// 		join
// 			Products p on td.ProductId = p.id
// 		group by
// 			p.ProductName
// 		order by
// 			TotalRevenue desc;
// 		`)
// 	if err != nil {
// 		log.Print("Error fetching records: ", err)
// 		return err
// 	}

// 	prodName := "Product Name"
// 	totalRevenue := "Total Revenue"

// 	fmt.Printf("%-24s %-12s\n", prodName, totalRevenue)
// 	fmt.Println("--------------------------------------")

// 	defer rows.Close()

// 	for rows.Next() {
// 		var productName string
// 		var totalRevenue float64

// 		err = rows.Scan(&productName, &totalRevenue)
// 		if err != nil {
// 			log.Print("Error scanning record: ", err)
// 			return err
// 		}

// 		fmt.Printf("%-24s %-12.2f\n", productName, totalRevenue)
// 	}

// 	return nil
// }
