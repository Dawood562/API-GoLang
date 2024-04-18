package database

import (
	"database/sql"
	"errors"
	"go-api/structs"
)

func GetItem(id string) (structs.Xyz, error) {
	rows := Db.QueryRow("SELECT * FROM things WHERE id = ?;", id)
	var thing structs.Xyz

	if err := rows.Scan(&thing.Id, &thing.Title, &thing.Description, &thing.Number, &thing.SomeBoolean); err != nil {
		if err == sql.ErrNoRows {
			return thing, errors.New("Item not found.")
		}
		return thing, errors.New("Error with turning records into struct.")
	}
	return thing, nil
}
