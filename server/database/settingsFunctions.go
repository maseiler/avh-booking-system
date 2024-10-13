package database

import (
	data "github.com/maseiler/avh-booking-system/server/data"
)

// GetSettings returns Array of settings from database
func GetSettings() []data.Setting {
	settingMap := []data.Setting{}
	rows, err := db.Query("SELECT * FROM settings;")
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		setting := data.Setting{}
		err := rows.Scan(&setting.Name, &setting.Value)
		settingMap = append(settingMap, setting)
		HandleDatabaseError(err)
	}
	err = rows.Err()
	HandleDatabaseError(err)
	return settingMap
}

// UpdateSetting updates setting values into database
func UpdateSetting(setting data.Setting) {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	// stmt, err := tx.Prepare("UPDATE settings SET value = ? WHERE settings.name = ?;")
	stmt, err := tx.Prepare("REPLACE INTO settings (name, value) VALUES(?, ?);")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(setting.Name, setting.Value)
	HandleDatabaseError(err)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
}
