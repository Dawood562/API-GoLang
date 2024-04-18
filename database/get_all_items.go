package database

import (
	"go-api/structs"

	"github.com/gin-gonic/gin"
)

func GetAllItems() ([]structs.Xyz, gin.H) {
	var things []structs.Xyz

	rows, err := Db.Query("SELECT * FROM things;")

	if err != nil {
		return nil, gin.H{"message": "Error with accessing the database."}
	}
	defer rows.Close()

	for rows.Next() {
		var thing structs.Xyz
		if err := rows.Scan(&thing.Id, &thing.Title, &thing.Description, &thing.Number, &thing.SomeBoolean); err != nil {
			return nil, gin.H{"message": "Error with turning database result into xyz struct."}
		}
		things = append(things, thing)
	}

	if err := rows.Err(); err != nil {
		return nil, gin.H{"message": "I haven't a scooby."}
	}
	return things, nil
}
