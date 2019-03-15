package database

import (
	"database/sql"
	"fmt"
	"strconv"

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

// getUsersByQuery returns list of users as requested in string
func getUsersByQuery(query string) []data.User {
	users := []data.User{}
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		user := data.User{}
		err := rows.Scan(&user.UserID, &user.BierName, &user.FirstName, &user.LastName, &user.Status, &user.Email, &user.Balance, &user.Phone, &user.MaxDebt)
		users = append(users, user)
		HandleDatabaseError(err)
		info := fmt.Sprintf("UserID: %d\nBierName: %s\nFirstName: %s\nLastName: %s\nStatus: %s\nEmail: %s\nPhone: %s\nBalance: %.2f\nMaxDebt: %d\n", user.UserID, user.BierName, user.FirstName, user.LastName, user.Status, user.Email, user.Phone, user.Balance, user.MaxDebt)
		fmt.Println(info)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	return users
}

// GetUsersOfColumnWithValue gets users which have a certain value in their column
func GetUsersOfColumnWithValue(column string, value string) []data.User {
	queryString := ""
	var err error
	if column == "BierName" || column == "FirstName" || column == "LastName" || column == "Status" || column == "Email" || column == "PhoneNumber" {
		queryString = fmt.Sprintf("SELECT * FROM users WHERE %s = \"%s\";", column, value)
	} else if column == "UserId" || column == "MaxDebt" {
		var intValue int
		intValue, err = strconv.Atoi(value)
		queryString = fmt.Sprintf("SELECT * FROM users WHERE %s = %d;", column, intValue)
	} else if column == "Balance" {
		var floatValue float64
		floatValue, err = strconv.ParseFloat(value, 32)
		queryString = fmt.Sprintf("SELECT * FROM users WHERE %s = %f;", column, floatValue)
	} else {
		panic("Invalid column")
	}
	HandleDatabaseError(err)
	return getUsersByQuery(queryString)
}

// NewUserExists returns true if user exists in database (based on BierName, FirstName and LastName)
func NewUserExists(newUser data.User) bool {
	queryString := fmt.Sprintf("SELECT * FROM users WHERE BierName = \"%s\" AND FirstName = \"%s\" AND LastName = \"%s\" AND Status = \"%s\";", newUser.BierName, newUser.FirstName, newUser.LastName, newUser.Status)
	users := getUsersByQuery(queryString)
	if len(users) == 0 {
		return false
	}
	return true
}

// UserExists returns true if user with same UserID exists in database
func UserExists(user data.User) bool {
	queryString := fmt.Sprintf("SELECT * FROM users WHERE UserId = %d;", user.UserID)
	users := getUsersByQuery(queryString)
	if len(users) == 0 {
		return false
	}
	return true
}

// AddUser adds a new user to database and prints info
func AddUser(newUser data.User) {
	// todo: get info from input
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO users(BierName, FirstName, LastName, Status, Email, Balance, PhoneNumber, MaxDebt) VAlUES(?, ?, ?, ?, ?, ?, ?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(newUser.BierName, newUser.FirstName, newUser.LastName, newUser.Status, newUser.Email, newUser.Balance, newUser.Phone, newUser.MaxDebt)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	stmt.Close()
}

// ModifyUser replaces all values of a user
func ModifyUser(user data.User) {
	queryString := fmt.Sprintf("UPDATE users SET BierName = \"%s\", FirstName = \"%s\", LastName = \"%s\", Status = \"%s\", Email = \"%s\", Balance = %f, PhoneNumber = \"%s\", MaxDebt = %d WHERE UserId = %d;", user.BierName, user.FirstName, user.LastName, user.Status, user.Email, user.Balance, user.Phone, user.MaxDebt, user.UserID)
	rows, err := db.Query(queryString)
	HandleDatabaseError(err)
	fmt.Println(rows)
}

// DeleteUser deletes a user with corresponding ID from database
func DeleteUser(user data.User) {
	queryString := fmt.Sprintf("DELETE FROM users WHERE UserId = %d;", user.UserID)
	rows, err := db.Query(queryString)
	HandleDatabaseError(err)
	fmt.Println(rows)
}

// GetItems returns all items except the first from database and prints them
// First Item is reserved
func GetItems() []data.Item {
	items := []data.Item{}
	queryString := "SELECT * FROM items WHERE ItemId > 1;"
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

// NewItemExists returns true if item exists in database (based on Name and Size)
func NewItemExists(newItem data.Item) bool {
	queryString := fmt.Sprintf("SELECT * FROM items WHERE Name = \"%s\" AND Size = %.2f;", newItem.Name, newItem.Size)
	items := getItemsByQuery(queryString)
	if len(items) == 0 {
		return false
	}
	return true
}

// ItemExists returns true if item with same ItemID exists in database
func ItemExists(item data.Item) bool {
	queryString := fmt.Sprintf("SELECT * FROM items WHERE ItemId = %d;", item.ItemID)
	items := getItemsByQuery(queryString)
	if len(items) == 0 {
		return false
	}
	return true
}

// getItemsByQuery returns list of items as requested in string
func getItemsByQuery(query string) []data.Item {
	items := []data.Item{}
	rows, err := db.Query(query)
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

// AddItem adds a new user to database and prints info
func AddItem(newItem data.Item) {
	// todo: get info from input
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO items(Name, Price, Size, Unit, Type) VAlUES(?, ?, ?, ?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(newItem.Name, newItem.Price, newItem.Size, newItem.Unit, newItem.Type)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	stmt.Close()
}

// ModifyItem replaces all values of a item
func ModifyItem(item data.Item) {
	queryString := fmt.Sprintf("UPDATE items SET Name = \"%s\", Price = \"%f\", Size = \"%f\", Unit = \"%s\", Type = \"%s\" WHERE ItemId = %d;", item.Name, item.Price, item.Size, item.Unit, item.Type, item.ItemID)
	rows, err := db.Query(queryString)
	HandleDatabaseError(err)
	fmt.Println(rows)
}

// DeleteItem deletes a item with corresponding ID from database
func DeleteItem(item data.Item) {
	queryString := fmt.Sprintf("DELETE FROM items WHERE ItemId = %d;", item.ItemID)
	rows, err := db.Query(queryString)
	HandleDatabaseError(err)
	fmt.Println(rows)
}
