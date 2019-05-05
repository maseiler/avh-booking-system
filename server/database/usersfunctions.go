package database

import (
	"fmt"
	"strconv"

	data "../data"
)

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
		// info := fmt.Sprintf("UserID: %d\nBierName: %s\nFirstName: %s\nLastName: %s\nStatus: %s\nEmail: %s\nPhone: %s\nBalance: %.2f\nMaxDebt: %d\n", user.UserID, user.BierName, user.FirstName, user.LastName, user.Status, user.Email, user.Phone, user.Balance, user.MaxDebt)
		// fmt.Println(info)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	fmt.Printf("Performed user query: \"%s\"\n", query)
	return users
}

// GetUsersOfColumnWithValue returns all users where value matches in specific column
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
	DeleteUserFromFavoriteItems(user)
}
