package endpoints

import (
	"go-api/database"
	"go-api/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetItems(context *gin.Context) {
	var things []structs.Xyz

	things, err := database.GetAllItems()

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	context.IndentedJSON(http.StatusOK, things)

}
