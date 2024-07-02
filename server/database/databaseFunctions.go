package database

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

// CreateDatabase creates database and all necessary tables
func CreateDatabase() {
	loginInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/", os.Getenv("AVHBS_DB_USER"), os.Getenv("AVHBS_DB_PASS"), os.Getenv("AVHBS_DB_IP"), os.Getenv("AVHBS_DB_PORT"))
	fmt.Println("LoginInfo:")
	fmt.Println(loginInfo)
	var err error
	db, err = sql.Open("mysql", loginInfo)
	HandleDatabaseError(err)

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", os.Getenv("AVHBS_DB_NAME"))
	_, err = db.Exec(query)
	HandleDatabaseError(err)

	db.Close()
	loginInfo = loginInfo + os.Getenv("AVHBS_DB_NAME") + "?parseTime=true"
	db, err = sql.Open("mysql", loginInfo)
	HandleDatabaseError(err)

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
		bier_name VARCHAR(50),
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		boat_name VARCHAR(50),
		status VARCHAR(20),
		email VARCHAR(50),
		phone VARCHAR(30),
		balance DECIMAL(6,2),
		max_debt INT
	);`
	_, err = db.Exec(createUsersTable)
	HandleDatabaseError(err)

	createItemsTable := `
	CREATE TABLE IF NOT EXISTS items(
		id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(50),
		type VARCHAR(20),
		size DECIMAL(6,2),
		unit VARCHAR(10),
		price DECIMAL(6,2),
		enabled BOOLEAN
	);`
	_, err = db.Exec(createItemsTable)
	HandleDatabaseError(err)

	// Altering because of a new Column
	alterItemsTable := `
	ALTER TABLE items ADD COLUMN IF NOT EXISTS enabled BOOLEAN NOT NULL DEFAULT '1' AFTER price
	;`
	_, err = db.Exec(alterItemsTable)
	HandleDatabaseError(err)

	createFavoriteItemsTable := `
	CREATE TABLE IF NOT EXISTS favorite_items (
		user_id INT NOT NULL,
		item_id INT NOT NULL,
		count INT,
		PRIMARY KEY (user_id, item_id)
	);`
	_, err = db.Exec(createFavoriteItemsTable)
	HandleDatabaseError(err)

	createBookingsTable := `
	CREATE TABLE IF NOT EXISTS bookings(
		id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
		time_stamp DATETIME NOT NULL,
		user_id INT NOT NULL,
		item_id INT NOT NULL,
		amount INT NOT NULL,
		total_price DECIMAL(6,2) NOT NULL,
		comment VARCHAR(255) NOT NULL,
		payment_method VARCHAR(10) NOT NULL
	);`
	_, err = db.Exec(createBookingsTable)
	HandleDatabaseError(err)

	createFeedbackTable := `
	CREATE TABLE IF NOT EXISTS feedback(
		id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
		time_stamp DATETIME,
		text VARCHAR(2000),
		name VARCHAR(20)
	);`
	_, err = db.Exec(createFeedbackTable)
	HandleDatabaseError(err)

	createPasswordsTable := `
	CREATE TABLE IF NOT EXISTS passwords(
		password VARCHAR(30)
	);`
	_, err = db.Exec(createPasswordsTable)
	HandleDatabaseError(err)

	if !passwordExists(db) {
		_, err = db.Exec("INSERT INTO passwords VALUES('admin');")
		HandleDatabaseError(err)
	}

	createSettingsTable := `
	CREATE TABLE IF NOT EXISTS settings(
		name VARCHAR(50) NOT NULL PRIMARY KEY,
		value VARCHAR(254) NOT NULL
	);`
	_, err = db.Exec(createSettingsTable)
	HandleDatabaseError(err)

	if !settingExists(db, "StripeAPIKey") {
		_, err = db.Exec("INSERT INTO settings VALUES('StripeAPIKey', '');")
		HandleDatabaseError(err)
	}
	// INSERT INTO `settings` (`name`, `value`) VALUES ('EMail-Host', 'mail.gmx.net');

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	newLoginInfo := fmt.Sprintf("%s:***@tcp(%s:%s)/%s", os.Getenv("AVHBS_DB_USER"), os.Getenv("AVHBS_DB_IP"), os.Getenv("AVHBS_DB_PORT"), os.Getenv("AVHBS_DB_NAME"))
	fmt.Println("AVHBS_DB_NAME:", os.Getenv("AVHBS_DB_NAME"))
	fmt.Printf("Database set up complete: %s\n-> %s\n", version, newLoginInfo)
}

func passwordExists(db *sql.DB) bool {
	var pws []string
	rows, err := db.Query("SELECT * FROM passwords;")
	HandleDatabaseError(err)
	defer rows.Close()
	if rows.Next() {
		var pw string
		err := rows.Scan(&pw)
		pws = append(pws, pw)
		HandleDatabaseError(err)
	}
	if len(pws) == 1 {
		return true
	}
	return false
}

func settingExists(db *sql.DB, name string) bool {
	query := fmt.Sprintf("SELECT * FROM settings WHERE name = '%s';", name)
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	fmt.Println(rows)
	var set []string
	if rows.Next() {
		set = append(set, "exists")
		HandleDatabaseError(err)
	}
	if len(set) == 1 {
		return true
	}
	return false
}
