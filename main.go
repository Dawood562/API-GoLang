package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/* Determine Data Structure */
type xyz struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Number      int    `json:"number"`
	SomeBoolean bool   `json:"someBoolean"`
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
	context.IndentedJSON(http.StatusOK, items)
}

func main() {
	// Run a server
	router := gin.Default()

	// Set an endpoint to have a GET request
	router.GET("/items", getItems)

	// Run a server
	router.Run("localhost:9000")
}
