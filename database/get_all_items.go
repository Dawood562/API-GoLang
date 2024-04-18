package database

import (
	"errors"
	"go-api/structs"
)

func GetAllItems() ([]structs.Xyz, error) {
	var things []structs.Xyz

	rows, err := Db.Query("SELECT * FROM things;")

	if err != nil {
		return nil, errors.New("Error with accessing the database.")
	}
	defer rows.Close()

	for rows.Next() {
		var thing structs.Xyz
		if err := rows.Scan(&thing.Id, &thing.Title, &thing.Description, &thing.Number, &thing.SomeBoolean); err != nil {
			return nil, errors.New("Error with turning database result into xyz struct.")
		}
		things = append(things, thing)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.New("I haven't a scooby.")
	}
	return things, nil
}
