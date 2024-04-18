package endpoints

import (
	"go-api/database"
	"go-api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateItem(context *gin.Context) {
	// Get item as given by the user
	var newItem structs.XyzPatch
	err := context.BindJSON(&newItem)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Make sure your data is formatted correctly."})
		return
	}

	// Convert it to a map to iterate through the values
	newMap, err := structs.Convert(newItem)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Make sure your data is formatted correctly."})
		return
	}
	// Check to make sure that all the values aren't nil
	var allNil bool = true
	for k := range newMap {
		if newMap[k] != nil {
			allNil = false
		}
	}
	if allNil {
		context.IndentedJSON(http.StatusBadRequest, newMap)
		return
	}

	// call update with ID and newmap with only non-nil values in it
	id := context.Param("id")
	rows, err := database.UpdateItem(id, newMap)
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": strconv.FormatInt(rows, 10) + " rows affected"})
}
