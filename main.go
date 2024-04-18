package main

import (
	"database/sql"
	"fmt"
	"go-api/endpoints"
	"go-api/structs"
	"log"
	"net/http"

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

	// Run a server
	router := gin.Default()

	// Set an endpoint to have a GET request
	router.GET("/items", endpoints.GetItems)

	// GET endpoint, takes an argument
	router.GET("/item/:id", endpoints.GetItemById)

	// POST endpoint
	router.POST("/addItem", addItem)

	// PATCH endpoint
	router.PATCH("/item/:id", endpoints.UpdateItem)

	// Run a server
	router.Run("localhost:9000")
}

func addItem(context *gin.Context) {
	var newItem structs.Xyz
	err := context.BindJSON(&newItem)
	if err != nil {
		return
	}

	result, err := db.Exec("INSERT INTO things (id, title, description, number, someboolean) VALUES (?, ?, ?, ?, ?);", newItem.Id, newItem.Title, newItem.Description, newItem.Number, newItem.SomeBoolean)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Item with specified ID probably already exists."})
	}
	id, err := result.RowsAffected()
	if err != nil {
		return
	}
	fmt.Println(id) // Need to do something with it so that the program runs
	context.IndentedJSON(http.StatusCreated, newItem)
}
