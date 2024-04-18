package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"

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
	Title       *string
	Description *string
	Number      *int
	SomeBoolean *bool
}

func getItems(context *gin.Context) {
	var things []xyz

	rows, err := db.Query("SELECT * FROM things;")
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, "Error with accessing the database.")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var thing xyz
		if err := rows.Scan(&thing.Id, &thing.Title, &thing.Description, &thing.Number, &thing.SomeBoolean); err != nil {
			context.IndentedJSON(http.StatusInternalServerError, "Error with turning database result into xyz struct.")
			return
		}
		things = append(things, thing)
	}

	if err := rows.Err(); err != nil {
		context.IndentedJSON(http.StatusInternalServerError, "I haven't a scooby.")
		return
	}
	context.IndentedJSON(http.StatusOK, things)

}

// func addItem(context *gin.Context) {
// 	var newItem xyz
// 	err := context.BindJSON(&newItem)
// 	if err != nil {
// 		return
// 	}

// 	items = append(items, newItem)
// 	context.IndentedJSON(http.StatusCreated, newItem)
// }

func getItem(id string) (xyz, error) {
	rows := db.QueryRow("SELECT * FROM things WHERE id = ?;", id)
	var thing xyz

	fmt.Println(rows)
	if err := rows.Scan(&thing.Id, &thing.Title, &thing.Description, &thing.Number, &thing.SomeBoolean); err != nil {
		if err == sql.ErrNoRows {
			return thing, errors.New("item Not Found")
		}
		fmt.Println(err)
		return thing, errors.New("error with turning database result into xyz struct")
	}
	return thing, nil
}

func getItemById(context *gin.Context) {
	id := context.Param("id")
	fmt.Println(id)
	it, err := getItem(id)
	if err != nil {
		fmt.Println(err)
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found."})
		return
	}

	context.IndentedJSON(http.StatusOK, it)
}

func getField(v *xyz, field string) int {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
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
	// router.POST("/addItem", addItem)

	// Run a server
	router.Run("localhost:9000")
}
