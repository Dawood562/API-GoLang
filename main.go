package main

import (
	"database/sql"
	"fmt"
	"go-api/endpoints"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Capture connection properties
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "things",
	}

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")

	// 'Create' server
	router := gin.Default()

	// Set an endpoint to have a GET request
	router.GET("/items", endpoints.GetItems)

	// GET endpoint, takes an argument
	router.GET("/item/:id", endpoints.GetItemById)

	// POST endpoint
	router.POST("/addItem", endpoints.AddItem)

	// PATCH endpoint
	router.PATCH("/item/:id", endpoints.UpdateItem)

	// Run a server
	router.Run("localhost:9000")
}
