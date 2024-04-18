package database

import (
	"database/sql"
	"go-api/structs"

	"github.com/gin-gonic/gin"
)

func GetItem(id string) (structs.Xyz, gin.H) {
	rows := Db.QueryRow("SELECT * FROM things WHERE id = ?;", id)
	var thing structs.Xyz

	if err := rows.Scan(&thing.Id, &thing.Title, &thing.Description, &thing.Number, &thing.SomeBoolean); err != nil {
		if err == sql.ErrNoRows {
			return thing, gin.H{"message": "Item not found."}
		}
		return thing, gin.H{"message": "Error with turning records into struct."}
	}
	return thing, nil
}
