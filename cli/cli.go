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
		fmt.Println("1. Add item")
		fmt.Println("2. Show all products")
		fmt.Println("3. Delete item")
		fmt.Println("4. Update item")
		fmt.Println("5. show report")

		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:

		case 2:

		case 3:
			c.deleteProduct()
		case 4:

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

				case 2:
					c.totalRevenueReport()
					return
				case 3:

				default:
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
