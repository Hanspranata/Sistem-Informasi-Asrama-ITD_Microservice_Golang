package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %s", err)
	}

	fmt.Println("Connected to the database!")
}
