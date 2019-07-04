package database

import (
	"fmt"
	"strconv"

	data "../data"
)

// getUsersByQuery returns list of users as requested in string
func getUsersByQuery(query string) []data.User {
	var users []data.User
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		user := data.User{}
		err := rows.Scan(&user.ID, &user.BierName, &user.FirstName, &user.LastName, &user.BoatName, &user.Status, &user.Email, &user.Phone, &user.Balance, &user.MaxDebt)
		users = append(users, user)
		HandleDatabaseError(err)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	// fmt.Printf("Performed user query: \"%s\"\n", query)
	return users
}

// GetUsersOfColumnWithValue returns all users where value matches in specific column
func GetUsersOfColumnWithValue(column string, value string) []data.User {
	queryString := ""
	var err error
	if column == "bier_name" || column == "first_name" || column == "last_name" || column == "boat_name" || column == "status" || column == "email" || column == "phone" {
		queryString = fmt.Sprintf("SELECT * FROM users WHERE %s = \"%s\";", column, value)
	} else if column == "id" || column == "max_debt" {
		var intValue int
		intValue, err = strconv.Atoi(value)
		queryString = fmt.Sprintf("SELECT * FROM users WHERE %s = %d;", column, intValue)
	} else if column == "balance" {
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
	queryString := fmt.Sprintf("SELECT * FROM users WHERE bier_name = \"%s\" AND first_name = \"%s\" AND last_name = \"%s\" AND status = \"%s\";", newUser.BierName, newUser.FirstName, newUser.LastName, newUser.Status)
	users := getUsersByQuery(queryString)
	if len(users) == 0 {
		return false
	}
	return true
}

// UserExists returns true if user with same UserID exists in database
func UserExists(user data.User) bool {
	queryString := fmt.Sprintf("SELECT * FROM users WHERE id = %d;", user.ID)
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
	stmt, err := tx.Prepare("INSERT INTO users(bier_name, first_name, last_name, boat_name, status, email, phone, balance, max_debt) VAlUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(newUser.BierName, newUser.FirstName, newUser.LastName, newUser.BoatName, newUser.Status, newUser.Email, newUser.Phone, newUser.Balance, newUser.MaxDebt)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	stmt.Close()
}

// ModifyUser replaces all values of a user
func ModifyUser(user data.User) {
	queryString := fmt.Sprintf("UPDATE users SET bier_name = \"%s\", first_name = \"%s\", last_name = \"%s\", status = \"%s\", email = \"%s\", balance = %f, phone = \"%s\", max_debt = %d, boat_name = \"%s\" WHERE id = %d", user.BierName, user.FirstName, user.LastName, user.Status, user.Email, user.Balance, user.Phone, user.MaxDebt, user.BoatName, user.ID)
	rows, err := db.Query(queryString)
	HandleDatabaseError(err)
	fmt.Println(rows)
}

// DeleteUser deletes a user with corresponding ID from database
func DeleteUser(user data.User) {
	queryString := fmt.Sprintf("DELETE FROM users WHERE id = %d;", user.ID)
	_, err := db.Query(queryString)
	HandleDatabaseError(err)
	// fmt.Println(rows)
	DeleteUserFromFavoriteItems(user)
}
