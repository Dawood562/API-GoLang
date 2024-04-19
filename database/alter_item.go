package database

import (
	"fmt"
	"reflect"
	"strings"
)

func UpdateItem(id string, givenMap map[string]interface{}) (int64, error) {
	// Check the non-nil values
	var toDelete []string
	for k, v := range givenMap {
		if reflect.ValueOf(v).IsNil() {
			toDelete = append(toDelete, k)
		}
	}
	for v := range toDelete {
		delete(givenMap, toDelete[v])
	}

	var updateString string = "UPDATE IGNORE things SET "
	// Convert the map into an array to be provided to the database
	values := make([]interface{}, 0, len(givenMap))
	for k, v := range givenMap {
		if !reflect.ValueOf(v).IsNil() {
			updateString += k + "=? "
			toDelete = append(toDelete, k)
		}
		switch t := givenMap[k].(type) {
		case *int:
			if givenMap[k].(*int) != nil {
				values = append(values, *givenMap[k].(*int))
			}
		case *float64:
			if givenMap[k].(*float64) != nil {
				values = append(values, *givenMap[k].(*float64))
			}
			fmt.Printf("%v", t)
		case *string:
			if givenMap[k].(*string) != nil {
				values = append(values, *givenMap[k].(*string))
			}
		case *bool:
			if givenMap[k].(*bool) != nil {
				values = append(values, *givenMap[k].(*bool))
			}
		default:
			fmt.Println("This type doesn't have a case switch.")
			fmt.Printf("Value %v, type %T", givenMap[k], givenMap[k])
		}
	}

	updateString = strings.Replace(updateString, "?", "?,", strings.Count(updateString, "?")-1)
	updateString += "WHERE id=?;"

	values = append(values, id)
	fmt.Println(updateString)
	fmt.Println(values)
	res, err := Db.Exec(updateString, values...)
	if err != nil {
		fmt.Println("At prepare")
		return 0, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Println("At rows affected")
		return rowsAffected, err
	}

	return rowsAffected, nil
}
