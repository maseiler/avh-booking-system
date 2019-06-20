package database

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
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	return pws
}

// TODO modify
