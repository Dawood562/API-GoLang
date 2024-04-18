package endpoints

import (
	"go-api/database"
	"go-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddItem(context *gin.Context) {
	var newItem structs.Xyz
	err := context.BindJSON(&newItem)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Make sure your data is formatted correctly."})
	}

	err = database.Insert(newItem)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
	}
	context.IndentedJSON(http.StatusCreated, newItem)
}
