package database

import (
	"errors"
	"go-api/structs"
)

func Insert(newItem structs.Xyz) error {
	result, err := Db.Exec("INSERT INTO things (id, title, description, number, someboolean) VALUES (?, ?, ?, ?, ?);", newItem.Id, newItem.Title, newItem.Description, newItem.Number, newItem.SomeBoolean)
	if err != nil {
		return errors.New("Item with specified ID probably already exists.")
	}
	_ = result
	return nil
}
