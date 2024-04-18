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
			return thing, errors.New("item not found")
		}
		return thing, errors.New("error with turning records into struct")
	}
	return thing, nil
}
