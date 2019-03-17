package database

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB

// TODO CreateDatabase()

// ConnectDatabase connects to the database and prints the version
func ConnectDatabase() {
	dbLogin := GetDatabaseLoginFromFile()
	loginInfo := fmt.Sprintf("%s:%s@/%s", dbLogin[0], dbLogin[1], dbLogin[2])
	var err error
	db, err = sql.Open("mysql", loginInfo)
	HandleDatabaseError(err)
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

// Checkout TODO
func Checkout(w http.ResponseWriter, r *http.Request) {
	//sql query
	fmt.Fprintf(w, "Check this out")
}
