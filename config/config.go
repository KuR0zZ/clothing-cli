package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// ConnectDB connects to the database
func ConnectDB() *sql.DB {
	var err error
	db_name := "postgres"
	db_user := "postgres.utunvkebldjwgbuwxhaa"
	db_pass := "adminklewear123"
	db_host := "aws-0-ap-southeast-1.pooler.supabase.com"

	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=require", db_name, db_user, db_pass, db_host))

	//DB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_golang1")
	if err != nil {
		log.Print("Error connecting to the database: ", err)
		log.Fatal(err)
	}

	// Check the connection
	if err = db.Ping(); err != nil {
		log.Print("Error pinging the database: ", err)
		log.Fatal(err)
	}

	log.Print("Connected to the database")

	return db
}
