package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// handler package import
	"github.com/emrctnn11/sezzle-payslice-backend/handlers"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// DB is the global database connection pool
var db *sql.DB

func main() {
	var err error

	// 1. Connect to MYSQL
	// Format: username:password@tcp(host:port)/dbname
	dsn := "sezzle_user:sezzle_password@tcp(localhost:3306)/product_db"

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening MySQL: ", err)
	}

	// Ping to verify
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to MySQL. Is Docker Running? ", err)
	}
	fmt.Println("Connected to MYSQL(PRODUCT CATALOG)")

	ph := &handlers.ProductHandler{DB: db}

	// defining routes
	http.HandleFunc("/products", ph.GetProducts)

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// API Handlers:
