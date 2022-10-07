package database

import (
	"fmt"
	"time"

	data "github.com/maseiler/avh-booking-system/server/data"
)

func AddNewUser(newUser data.User) {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	dtNow := time.Now()
	stmt, err := tx.Prepare("INSERT INTO users(creation, client, bier_name, first_name, last_name, boat_name, status, email, phone, balance, max_debt) VAlUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(dtNow, newUser.Client, newUser.BierName, newUser.FirstName, newUser.LastName, newUser.BoatName, newUser.Status, newUser.Email, newUser.Phone, newUser.Balance, newUser.MaxDebt)
	HandleDatabaseError(err)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
}

// Returns a list of users as requested in string
func GetUsersByQuery(query string) []data.User {
	var users []data.User
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		user := data.User{}
		err := rows.Scan(&user.Id, &user.Creation, &user.Client, &user.BierName, &user.FirstName, &user.LastName,
			&user.BoatName, &user.Status, &user.Email, &user.Phone, &user.Balance, &user.MaxDebt)
		users = append(users, user)
		HandleDatabaseError(err)
	}
	err = rows.Err()
	HandleDatabaseError(err)
	// fmt.Printf("Performed user query: \"%s\"\n", query)
	return users
}

func NewUserExists(newUser data.User) bool {
	queryString := fmt.Sprintf("SELECT * FROM users WHERE bier_name = \"%s\" AND first_name = \"%s\" AND last_name = \"%s\" AND status = \"%s\";", newUser.BierName, newUser.FirstName, newUser.LastName, newUser.Status)
	users := GetUsersByQuery(queryString)
	return len(users) != 0
}
