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

// Init variables
var items = []xyz{
	{Id: "Unique Identifier #1", Title: "Title 1", Description: "Here is some more information about title.", Number: 1, SomeBoolean: true},
	{Id: "mov_inception", Title: "Inception", Description: "This film is about drugs and draeming.", Number: 56, SomeBoolean: false},
	{Id: "tv_rickandmorty", Title: "Rick & Morty", Description: "I turned myself into a pickle, Morty! I'm Pickle Riiiiiiick!", Number: 6786789, SomeBoolean: false},
	{Id: "docu_totaltrust", Title: "Total Trust", Description: "A documentary about surveillance and censorship in China.", Number: 562, SomeBoolean: true},
	{Id: "soft_vscode", Title: "Visual Studio Code", Description: "IDE.", Number: 3589, SomeBoolean: false},
	{Id: "web_google", Title: "Google", Description: "A well-known search engine.", Number: 99, SomeBoolean: true},
}

func getItems(context *gin.Context) {
	var things []xyz

	rows, err := db.Query("SELECT * FROM things;")
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, "Error with accessing the database.")
		return
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
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

func addItem(context *gin.Context) {
	var newItem xyz
	err := context.BindJSON(&newItem)
	if err != nil {
		return
	}

	items = append(items, newItem)
	context.IndentedJSON(http.StatusCreated, newItem)
}

func getItem(id string) (*xyz, error) {
	for i, t := range items {
		if t.Id == id {
			return &items[i], nil
		}
	}
	return nil, errors.New("Todo Not Found")
}

func getItemById(context *gin.Context) {
	id := context.Param("id")
	it, err := getItem(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Item not found."})
		return
	}

	context.IndentedJSON(http.StatusOK, it)
	return
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

	// POST endpoint
	router.POST("/addItem", addItem)

	// GET endpoint, takes an argument
	router.GET("/item/:id", getItemById)

	// Run a server
	router.Run("localhost:9000")
}
