package database

import (
	"database/sql"
	"fmt"
	"os"
)

var db *sql.DB

// CreateDatabase creates database and all necessary tables, constraints, triggers and values
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

	// create tables with constraints
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
		max_debt INT NOT NULL CHECK (max_debt < 0),
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
			size DECIMAL(6,2) NOT NULL CHECK (size > 0),
			unit VARCHAR(10) NOT NULL,
			price DECIMAL(6,2) NOT NULL CHECK (price > 0),
			PRIMARY KEY (id)
		);`
	_, err = db.Exec(createItemsTable)
	HandleDatabaseError(err)

	createFavoriteItemsTable := `
		CREATE TABLE IF NOT EXISTS favorite_items (
			user_id INT NOT NULL,
			item_id INT NOT NULL,
			count INT NOT NULL CHECK (count >= 0),
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
			amount INT NOT NULL CHECK (amount > 0),
			price DECIMAL(6,2) NOT NULL CHECK (price < 0),
			comment VARCHAR(255),
			PRIMARY KEY (id)
		);`
	_, err = db.Exec(createBookingsTable)
	HandleDatabaseError(err)

	createPaymentsTable := `
		CREATE TABLE IF NOT EXISTS payments(
			id INT NOT NULL AUTO_INCREMENT,
			time_stamp DATETIME NOT NULL,
			client VARCHAR(20) NOT NULL,
			user_id INT NOT NULL,
			amount INT NOT NULL CHECK (amount > 0),
			payment_method VARCHAR(10) NOT NULL,
			comment VARCHAR(255),
			PRIMARY KEY (id)
		);`
	_, err = db.Exec(createPaymentsTable)
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
	_, err = db.Exec("ALTER TABLE favorite_items ADD CONSTRAINT FOREIGN KEY (user_id) REFERENCES users(id);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE favorite_items ADD CONSTRAINT FOREIGN KEY (item_id) REFERENCES items(id);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE bookings ADD CONSTRAINT FOREIGN KEY (client) REFERENCES clients(name);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE bookings ADD CONSTRAINT FOREIGN KEY (user_id) REFERENCES users(id);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE bookings ADD CONSTRAINT FOREIGN KEY (item_id) REFERENCES items(id);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE payments ADD CONSTRAINT FOREIGN KEY (client) REFERENCES clients(name);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE payments ADD CONSTRAINT FOREIGN KEY (user_id) REFERENCES users(id);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE feedback ADD CONSTRAINT FOREIGN KEY (client) REFERENCES clients(name);")
	HandleDatabaseError(err)
	_, err = db.Exec("ALTER TABLE passwords ADD CONSTRAINT FOREIGN KEY (client) REFERENCES clients(name);")
	HandleDatabaseError(err)

	// add triggers
	incrementFavoriteItemTrigger := `
	CREATE TRIGGER increment_favorite_item
	AFTER INSERT ON bookings
	FOR EACH ROW
	BEGIN
		IF (NEW.item_id > 0) THEN
			IF (SELECT EXISTS(SELECT * FROM favorite_items WHERE favorite_items.user_id=NEW.user_id AND favorite_items.item_id=NEW.item_id)) = 1 THEN
				UPDATE favorite_items SET favorite_items.count = favorite_items.count + NEW.amount WHERE favorite_items.user_id = NEW.user_id AND favorite_items.item_id = NEW.item_id;
			ELSE
				INSERT INTO favorite_items VALUES(NEW.user_id, NEW.item_id, NEW.amount);
			END IF;
		END IF;
	END;
	`
	_, err = db.Exec(incrementFavoriteItemTrigger)
	HandleDatabaseError(err)

	updateBalanceTrigger := `
	CREATE TRIGGER update_balance
	AFTER INSERT ON payments
	FOR EACH ROW
	BEGIN
		UPDATE users SET users.balance = users.balance + NEW.amount WHERE users.id = NEW.user_id;
	END;
	`
	_, err = db.Exec(updateBalanceTrigger)
	HandleDatabaseError(err)

	// insert values
	_, err = db.Exec("INSERT IGNORE INTO clients VALUES('Server', NOW());")
	HandleDatabaseError(err)

	insertPasswordIfNotExists := `
	INSERT INTO passwords (password, creation, client)
	SELECT 'admin', NOW(), 'Server'
	WHERE NOT EXISTS (SELECT * FROM passwords);
	`
	_, err = db.Exec(insertPasswordIfNotExists)
	HandleDatabaseError(err)

	// reconnect
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	newLoginInfo := fmt.Sprintf("%s:***@tcp(%s:%s)/%s", os.Getenv("AVHBS_DB_USER"), os.Getenv("AVHBS_DB_IP"), os.Getenv("AVHBS_DB_PORT"), os.Getenv("AVHBS_DB_NAME"))
	fmt.Println("AVHBS_DB_NAME:", os.Getenv("AVHBS_DB_NAME"))
	fmt.Printf("Database set up complete: %s\n-> %s\n", version, newLoginInfo)
}
