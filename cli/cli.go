package cli

import (
	"clothing-cli/handler"
	"fmt"
	"log"
)

type CLI struct {
	Handler handler.Handler
}

func NewCLI(handler handler.Handler) *CLI {
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

	// show initial menu
	c.showMenu()
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
		fmt.Scan(&choice)

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
				fmt.Scan(&choice2)

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
				fmt.Println("Do you want to back to main menu? (y/n)")
				fmt.Scan(&cont)

				if cont == "y" {
					break
				} else if cont == "n" {
					return
				} else {
					fmt.Println("Invalid choice")
				}

			}

		case 6:
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func (c *CLI) addProduct() {
	var productName string
	var price float64
	var stock int

	fmt.Print("Enter product name: ")
	fmt.Scan(&productName)

	fmt.Print("Enter product price: ")
	fmt.Scan(&price)

	fmt.Print("Enter product stock: ")
	fmt.Scan(&stock)

	err := c.Handler.AddProduct(productName, price, stock)
	if err != nil {
		log.Print("Error adding product: ", err)
		log.Fatal(err)
	}

	fmt.Println("Successfully add new product")
}

func (c *CLI) updateProduct() {
	var productId int
	var productName string
	var price float64
	var stock int

	fmt.Print("Enter product id: ")
	fmt.Scan(&productId)

	fmt.Print("Enter product name: ")
	fmt.Scan(&productName)

	fmt.Print("Enter product price: ")
	fmt.Scan(&price)

	fmt.Print("Enter product stock: ")
	fmt.Scan(&stock)

	err := c.Handler.UpdateProduct(productId, productName, price, stock)
	if err != nil {
		log.Print("Error updating product: ", err)
		log.Fatal(err)
	}

	fmt.Println("Successfully update a product")
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
		log.Print("Error showing Total Revenue Per Game Report: ", err)
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
