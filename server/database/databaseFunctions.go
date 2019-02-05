package database

import (
	"database/sql"
	"fmt"

	"../data"
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

// GetItems returns all items from database and prints them
func GetItems() []data.Item {
	items := []data.Item{}
	queryString := "SELECT * FROM items;"
	// queryString = fmt.Sprintf("SELECT * FROM items WHERE %s = '%s';", column, value)
	rows, err := db.Query(queryString)
	HandleDatabaseError(err)

	defer rows.Close()
	for rows.Next() {
		item := data.Item{}
		err := rows.Scan(&item.ItemID, &item.Name, &item.Price, &item.Size, &item.Unit, &item.Type)
		items = append(items, item)
		HandleDatabaseError(err)
		info := fmt.Sprintf("ItemID: %d\nName: %s\nPrice: %.2f\nSize: %.2f\nUnit: %s\nType: %s\n", item.ItemID, item.Name, item.Price, item.Size, item.Unit, item.Type)
		fmt.Println(info)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	return items
}

// AddItems adds a new item to database and prints info
// func AddItem(newItem dataP.Item) {
// 	tx, err := db.Begin()
// 	HandleDatabaseError(err)
// 	stmt, err := tx.Prepare("INSERT INTO items(Name, Price, Size, Unit) VAlUES(?, ?, ?, ?)")
// 	HandleTxError(*tx, err)
// 	defer stmt.Close()
// 	res, err := stmt.Exec(newItem.Name, newItem.Price, newItem.Size, newItem.Unit)
// 	TxRowsAffected(res, *tx)
// 	err = tx.Commit()
// 	HandleDatabaseError(err)
// 	stmt.Close()
// }

// GetUsers returns all users from database and prints them
// func GetUsers() []data.User {
// 	users := []data.User{}
// 	queryString := "SELECT * FROM users;"
// 	// queryString = fmt.Sprintf("SELECT * FROM users WHERE %s = '%s';", column, value)
// 	rows, err := db.Query(queryString)
// 	HandleDatabaseError(err)
// 	defer rows.Close()
// 	for rows.Next() {
// 		user := data.User{}
// 		err := rows.Scan(&user.UserID, &user.BierName, &user.FirstName, &user.LastName, &user.Status, &user.Email, &user.Balance, &user.Phone)
// 		users = append(users, user)
// 		HandleDatabaseError(err)
// 		info := fmt.Sprintf("UserID: %d\nBierName: %s\nFirstName: %s\nLastName: %s\nStatus: %s\nEmail: %s\nPhone: %s\nBalance: %.2f\n", user.UserID, user.BierName, user.FirstName, user.LastName, user.Status, user.Email, user.Phone, user.Balance)
// 		fmt.Println(info)
// 	}
// 	defer rows.Close()
// 	err = rows.Err()
// 	HandleDatabaseError(err)
// 	return users
// }

// GetUsersOfStatus returns Users from database depending on status or all users if string is empty
func GetUsersOfStatus(status string) []data.User {
	users := []data.User{}
	queryString := ""
	if status != "" {
		queryString = fmt.Sprintf("SELECT * FROM users WHERE Status = \"%s\";", status)
	} else {
		queryString = "SELECT * FROM users;"
	}
	rows, err := db.Query(queryString)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		user := data.User{}
		err := rows.Scan(&user.UserID, &user.BierName, &user.FirstName, &user.LastName, &user.Status, &user.Email, &user.Balance, &user.Phone)
		users = append(users, user)
		HandleDatabaseError(err)
		info := fmt.Sprintf("UserID: %d\nBierName: %s\nFirstName: %s\nLastName: %s\nStatus: %s\nEmail: %s\nPhone: %s\nBalance: %.2f\n", user.UserID, user.BierName, user.FirstName, user.LastName, user.Status, user.Email, user.Phone, user.Balance)
		fmt.Println(info)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	return users
}

// AddUser adds a new user to database and prints info
func AddUser(newUser data.User) {
	// todo: get info from input
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO users(BierName, FirstName, LastName, Status, Email, PhoneNumber, Balance) VAlUES(?, ?, ?, ?, ?, ?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(newUser.BierName, newUser.FirstName, newUser.LastName, newUser.Status, newUser.Email, newUser.Phone, 0)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	stmt.Close()
}
