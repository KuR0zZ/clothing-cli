// Go & MySQL CRUD Example
package main

import (
	"clothing-cli/cli"
	"clothing-cli/config"
	"clothing-cli/repository"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func main() {
	db := config.ConnectDB()

	handler := repository.NewRepo(db)

	cli := cli.NewCLI(handler)
	cli.Init()
}
