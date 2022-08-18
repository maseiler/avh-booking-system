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
	fmt.Println("Database Login Info:")
	fmt.Println(loginInfo)
	var err error
	db, err = sql.Open("mysql", loginInfo)
	HandleDatabaseError(err)

	createDatabaseQuery := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", os.Getenv("AVHBS_DB_NAME"))
	_, err = db.Exec(createDatabaseQuery)
	HandleDatabaseError(err)
	db.Close()

	loginInfo = loginInfo + os.Getenv("AVHBS_DB_NAME") + "?parseTime=true"
	db, err = sql.Open("mysql", loginInfo)
	HandleDatabaseError(err)

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INT NOT NULL AUTO_INCREMENT,
		creation DATETIME NOT NULL,
		client VARCHAR(20) NOT NULL,
		bier_name VARCHAR(50),
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		boat_name VARCHAR(50),
		status VARCHAR(20) NOT NULL,
		email VARCHAR(50),
		phone VARCHAR(30),
		balance DECIMAL(6,2) NOT NULL,
		max_debt INT NOT NULL,
		PRIMARY KEY (id)
	);`
	_, err = db.Exec(createUsersTable)
	HandleDatabaseError(err)

	createItemsTable := `
		CREATE TABLE IF NOT EXISTS items(
			id INT NOT NULL AUTO_INCREMENT,
			creation DATETIME NOT NULL,
			client VARCHAR(20) NOT NULL,
			name VARCHAR(50) NOT NULL,
			type VARCHAR(20) NOT NULL,
			size DECIMAL(6,2) NOT NULL,
			unit VARCHAR(10) NOT NULL,
			price DECIMAL(6,2) NOT NULL,
			PRIMARY KEY (id)
		);`
	_, err = db.Exec(createItemsTable)
	HandleDatabaseError(err)

	// TODO create trigger: increment, if entry in bookings
	createFavoriteItemsTable := `
		CREATE TABLE IF NOT EXISTS favorite_items (
			user_id INT NOT NULL,
			item_id INT NOT NULL,
			count INT NOT NULL,
			PRIMARY KEY (user_id, item_id)
		);`
	_, err = db.Exec(createFavoriteItemsTable)
	HandleDatabaseError(err)

	createBookingsTable := `
		CREATE TABLE IF NOT EXISTS bookings(
			id INT NOT NULL AUTO_INCREMENT,
			time_stamp DATETIME NOT NULL,
			client VARCHAR(20) NOT NULL,
			user_id INT NOT NULL,
			item_id INT NOT NULL,
			amount INT NOT NULL,
			price DECIMAL(6,2) NOT NULL,
			comment VARCHAR(255) NOT NULL,
			payment_method VARCHAR(10) NOT NULL,
			PRIMARY KEY (id)
		);`
	_, err = db.Exec(createBookingsTable)
	HandleDatabaseError(err)

	createFeedbackTable := `
		CREATE TABLE IF NOT EXISTS feedback(
			id INT NOT NULL AUTO_INCREMENT,
			time_stamp DATETIME NOT NULL,
			client VARCHAR(20) NOT NULL,
			text VARCHAR(2000),
			name VARCHAR(20),
			PRIMARY KEY (id)
		);`
	_, err = db.Exec(createFeedbackTable)
	HandleDatabaseError(err)

	createClientsTable := `
	CREATE TABLE IF NOT EXISTS clients(
		name VARCHAR(20) NOT NULL,
		creation DATETIME NOT NULL,
		PRIMARY KEY (name)
	);`
	_, err = db.Exec(createClientsTable)
	HandleDatabaseError(err)

	createPasswordsTable := `
		CREATE TABLE IF NOT EXISTS passwords(
			password VARCHAR(30) NOT NULL,
			creation DATETIME NOT NULL,
			client VARCHAR(20) NOT NULL,
			PRIMARY KEY (password)
		);`
	_, err = db.Exec(createPasswordsTable)
	HandleDatabaseError(err)

	// add foreign keys
	_, err = db.Exec("ALTER TABLE users ADD CONSTRAINT FOREIGN KEY (client) REFERENCES clients(name);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE items ADD CONSTRAINT FOREIGN KEY (client) REFERENCES clients(name);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE bookings ADD CONSTRAINT FOREIGN KEY (client) REFERENCES clients(name);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE feedback ADD CONSTRAINT FOREIGN KEY (client) REFERENCES clients(name);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE passwords ADD CONSTRAINT FOREIGN KEY (client) REFERENCES clients(name);")
	HandleDatabaseError(err)

	// insert values
	_, err = db.Exec("INSERT IGNORE INTO clients VALUES('Server', NOW());")
	HandleDatabaseError(err)

	if !passwordExists(db) {
		_, err = db.Exec("INSERT IGNORE INTO passwords VALUES('admin', NOW(), 'Server');")
		HandleDatabaseError(err)
	}

	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	newLoginInfo := fmt.Sprintf("%s:***@tcp(%s:%s)/%s", os.Getenv("AVHBS_DB_USER"), os.Getenv("AVHBS_DB_IP"), os.Getenv("AVHBS_DB_PORT"), os.Getenv("AVHBS_DB_NAME"))
	fmt.Println("AVHBS_DB_NAME:", os.Getenv("AVHBS_DB_NAME"))
	fmt.Printf("Database set up complete: %s\n-> %s\n", version, newLoginInfo)
}

func passwordExists(db *sql.DB) bool {
	var pws []string
	rows, err := db.Query("SELECT password FROM passwords;")
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
