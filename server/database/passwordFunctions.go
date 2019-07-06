package database

import "fmt"

// GetPasswords returns all passwords from database
func GetPasswords() []string {
	var pws []string
	rows, err := db.Query("SELECT * FROM passwords;")
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		pw := ""
		err := rows.Scan(&pw)
		pws = append(pws, pw)
		HandleDatabaseError(err)
	}
	err = rows.Err()
	HandleDatabaseError(err)
	return pws
}

// ModifyPassword updates matching password in database
func ModifyPassword(oldPw string, newPw string) {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	query := fmt.Sprintf("UPDATE passwords SET password = \"%s\" WHERE password = \"%s\";", newPw, oldPw)
	stmt, err := tx.Prepare(query)
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec()
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
}
