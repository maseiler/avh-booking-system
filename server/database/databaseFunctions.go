package database

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

// TODO CreateDatabase()

// ConnectDatabase connects to the database and prints the version
func ConnectDatabase() {
	loginInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("AVHBS_DB_USER"), os.Getenv("AVHBS_DB_PASS"), os.Getenv("AVHBS_DB_IP"), os.Getenv("AVHBS_DB_PORT"), os.Getenv("AVHBS_DB_NAME")+"?parseTime=true")
	fmt.Println(loginInfo)
	var err error
	db, err = sql.Open("mysql", loginInfo)
	HandleDatabaseError(err)
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}
