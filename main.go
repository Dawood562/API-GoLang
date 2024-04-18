package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Determine Data Structure
type xyz struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Number      int    `json:"number"`
	SomeBoolean bool   `json:"someBoolean"`
}

// Structure specifically for PATCH requests to ensure I can tell which data needs to be patched
type xyzPatch struct {
	Id          *string
	Title       *string
	Description *string
	Number      *int
	SomeBoolean *bool
}

func getItems(context *gin.Context) {
	var things []xyz

	rows, err := db.Query("SELECT * FROM things;")
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error with accessing the database."})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var thing xyz
		if err := rows.Scan(&thing.Id, &thing.Title, &thing.Description, &thing.Number, &thing.SomeBoolean); err != nil {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error with turning database result into xyz struct."})
			return
		}
		things = append(things, thing)
	}

	if err := rows.Err(); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "I haven't a scooby."})
		return
	}
	context.IndentedJSON(http.StatusOK, things)

}

func addItem(context *gin.Context) {
	var newItem xyz
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

func getItem(id string) (xyz, gin.H) {
	rows := db.QueryRow("SELECT * FROM things WHERE id = ?;", id)
	var thing xyz

	fmt.Println(rows)
	if err := rows.Scan(&thing.Id, &thing.Title, &thing.Description, &thing.Number, &thing.SomeBoolean); err != nil {
		if err == sql.ErrNoRows {
			return thing, gin.H{"message": "Item not found."}
		}
		fmt.Println(err)
		return thing, gin.H{"message": "Error with turning records into struct."}
	}
	return thing, nil
}

func getItemById(context *gin.Context) {
	id := context.Param("id")
	it, err := getItem(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, err)
		return
	}

	context.IndentedJSON(http.StatusOK, it)
}

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
	router.GET("/items", getItems)

	// GET endpoint, takes an argument
	router.GET("/item/:id", getItemById)

	// POST endpoint
	router.POST("/addItem", addItem)

	// Run a server
	router.Run("localhost:9000")
}
