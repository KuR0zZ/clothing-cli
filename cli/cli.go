package cli

import (
	"bufio"
	"clothing-cli/repository"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CLI struct {
	Handler repository.Repository
}

func NewCLI(handler repository.Repository) *CLI {
	return &CLI{
		Handler: handler,
	}
}

func (c *CLI) Init() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	// show user login
	c.userLogin()
}

func (c *CLI) showMenu() {
	for {

		fmt.Println("Welcome to KLE-WEAR Clothing Server")
		fmt.Println("Main Menu:")
		fmt.Println("1. Add item")
		fmt.Println("2. Show all products")
		fmt.Println("3. Delete item")
		fmt.Println("4. Update item")
		fmt.Println("5. show report")
		fmt.Println("6. Exit")

		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			c.addProduct()
		case 2:
			c.showAllProducts()
		case 3:
			c.deleteProduct()
		case 4:
			c.updateProduct()
		case 5:
			for {
				var choice2 int
				fmt.Println("Show report:")
				fmt.Println("1. Current Stock")
				fmt.Println("2. Total Revenue")
				fmt.Println("3. Transaction Report")
				fmt.Print("Enter your choice: ")
				fmt.Scanln(&choice2)

				switch choice2 {
				case 1:
					c.reportCurrentStock()
				case 2:
					c.totalRevenueReport()
				case 3:
					c.customersTransactionsReport()
				default:
					fmt.Println("Invalid choice")
				}

				// ask user if they want to continue
				var cont string
				fmt.Println("=====================================")
				fmt.Println("Do you want to back to main menu? (y/n)")
				fmt.Scanln(&cont)

				if cont == "y" {
					c.showMenu()
				} else if cont == "n" {
					return
				} else {
					fmt.Println("Invalid choice")
					return
				}

			}

		case 6:
			return
		default:
			fmt.Println("Invalid choice")
		}

		// ask user if they want to continue
		var cont string

		fmt.Println("=====================================")
		fmt.Println("Do you want to back to main menu? (y/n)")
		fmt.Scanln(&cont)

		if cont == "n" {
			return
		} else if cont != "y" {
			fmt.Println("Invalid choice")
			return
		} else {
			continue
		}
	}
}

func (c *CLI) userLogin() {
	var email, password string

	fmt.Print("Enter email: ")
	fmt.Scanln(&email)

	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	_, err := c.Handler.UserLogin(email, password)

	if err != nil {
		log.Print("Error: ", err)
		c.userLogin()
	} else {
		c.showMenu()
	}
}

func (c *CLI) addProduct() {
	var productName string
	var price float64
	var stock int
	var err error

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter product name: ")
		productName, _ = reader.ReadString('\n')
		productName = strings.TrimSpace(productName)

		if productName == "" {
			fmt.Println("No input provided. Please enter a product name")
			continue
		}

		break
	}

	for {
		fmt.Print("Enter product price: ")
		priceInput, _ := reader.ReadString('\n')
		priceInput = strings.TrimSpace(priceInput)

		price, err = strconv.ParseFloat(priceInput, 64)
		if err != nil {
			fmt.Println("Invalid price. Please enter a valid number")
			continue
		}

		break
	}

	for {
		fmt.Print("Enter product stock: ")
		stockInput, _ := reader.ReadString('\n')
		stockInput = strings.TrimSpace(stockInput)

		stock, err = strconv.Atoi(stockInput)
		if err != nil {
			fmt.Println("Invalid stock. Please enter a valid integer")
			continue
		}

		break
	}

	err = c.Handler.AddProduct(productName, price, stock)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully added new product")
}

func (c *CLI) updateProduct() {
	var productId int
	var productName string
	var price float64
	var stock int
	var err error

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter product id: ")
		productIdInput, _ := reader.ReadString('\n')
		productIdInput = strings.TrimSpace(productIdInput)

		productId, err = strconv.Atoi(productIdInput)

		if err != nil {
			fmt.Println("Invalid type for product id")
			continue
		}

		break
	}

	fmt.Print("Enter product name: ")
	productName, _ = reader.ReadString('\n')
	productName = strings.TrimSpace(productName)

	for {
		fmt.Print("Enter product price: ")
		priceInput, _ := reader.ReadString('\n')
		priceInput = strings.TrimSpace(priceInput)

		if priceInput == "" {
			break
		}

		price, err = strconv.ParseFloat(priceInput, 64)
		if err != nil {
			fmt.Println("Invalid price. Please enter a valid number")
			continue
		}

		break
	}

	for {
		fmt.Print("Enter product stock: ")
		stockInput, _ := reader.ReadString('\n')
		stockInput = strings.TrimSpace(stockInput)

		if stockInput == "" {
			break
		}

		stock, err = strconv.Atoi(stockInput)
		if err != nil {
			fmt.Println("Invalid stock. Please enter a valid integer")
			continue
		}

		break
	}

	err = c.Handler.UpdateProduct(productId, productName, price, stock)
	if err != nil {
		log.Fatal(err)
	}
}

func (c *CLI) customersTransactionsReport() {
	err := c.Handler.CustomersTransactionsReport()
	if err != nil {
		log.Print("Error showing customers transactions report")
		log.Fatal(err)
	}
}

func (c *CLI) deleteProduct() {
	var name string

	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	err := c.Handler.DeleteProduct(name)
	if err != nil {
		log.Print("Error deleting product: ", err)
		log.Fatal(err)
	}
}

func (c *CLI) totalRevenueReport() {

	err := c.Handler.TotalRevenueReport()
	if err != nil {
		log.Print("Error showing Total Revenue Report: ", err)
		log.Fatal(err)
	}

}

// create funtion error handling for showAllProducts
func (c *CLI) showAllProducts() {
	err := c.Handler.ShowAllProducts()
	if err != nil {
		log.Print("Error showing all products: ", err)
		log.Fatal(err)
	}
}

// create function error handling for reportCurrentStock
func (c *CLI) reportCurrentStock() {
	err := c.Handler.CurrentStockReport()
	if err != nil {
		log.Print("Error showing current stock: ", err)
		log.Fatal(err)
	}
}
