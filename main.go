// Go & MySQL CRUD Example
package main

import (
	"clothing-cli/cli"
	"clothing-cli/config"
	"clothing-cli/handler"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	db := config.ConnectDB()

	handler := handler.NewHandler(db)

	cli := cli.NewCLI(handler)
	cli.Init()
}
