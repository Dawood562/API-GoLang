package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {

	// Capture connection properties
	// Fetch the database auth credentials
	dbUsername, dbPassword, dbName := fetchDBAuth()
	cfg := mysql.Config{
		User:   dbUsername,
		Passwd: dbPassword,
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: dbName,
	}

	// Get a database handle
	var err error
	Db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := Db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected")
}

/*
Fetches database login details from .env
db_details.txt should be in database folder with following content structure:
<username>
<database name>
<password>
*/
func fetchDBAuth() (string, string, string) {
	// Username, Password, Database Name
	return "root", "root", "things"
}
