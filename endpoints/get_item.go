package endpoints

import (
	"go-api/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItemById(context *gin.Context) {
	id := context.Param("id")
	it, err := database.GetItem(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, err)
		return
	}

	context.IndentedJSON(http.StatusOK, it)
}
