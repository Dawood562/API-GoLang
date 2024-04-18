package database

import (
	"fmt"
	"strings"
)

func UpdateItem(id string, givenMap map[string]interface{}) (int64, error) {
	fmt.Println("-----------------------------------------------------")
	var updateString string = "UPDATE IGNORE things SET "
	// Check the non-nil values
	var toDelete []string
	for k := range givenMap {
		switch t := givenMap[k].(type) {
		case *int:
			if givenMap[k].(*int) == nil {
				toDelete = append(toDelete, k)
			}
		case *float64:
			if givenMap[k].(*float64) == nil {
				toDelete = append(toDelete, k)
			}
			fmt.Printf("%v", t)
		case *string:
			if givenMap[k].(*string) == nil {
				toDelete = append(toDelete, k)
			}
		case *bool:
			if givenMap[k].(*bool) == nil {
				toDelete = append(toDelete, k)
			}
		default:
			fmt.Println("What the fuck.")
			fmt.Printf("Value %v, type %T", givenMap[k], givenMap[k])
		}
	}
	for v := range toDelete {
		delete(givenMap, toDelete[v])
	}

	fmt.Println(updateString)
	// Convert the map into an array to be provided to the database
	values := make([]interface{}, 0, len(givenMap))
	for k := range givenMap {
		switch t := givenMap[k].(type) {
		case *int:
			if givenMap[k].(*int) != nil {
				updateString += k + "=? "
				values = append(values, *givenMap[k].(*int))
			}
		case *float64:
			if givenMap[k].(*float64) != nil {
				updateString += k + "=? "
				values = append(values, *givenMap[k].(*float64))
			}
			fmt.Printf("%v", t)
		case *string:
			if givenMap[k].(*string) != nil {
				updateString += k + "=? "
				values = append(values, *givenMap[k].(*string))
			}
		case *bool:
			if givenMap[k].(*bool) != nil {
				updateString += k + "=? "
				values = append(values, *givenMap[k].(*bool))
			}
		default:
			fmt.Println("What the fuck.")
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
