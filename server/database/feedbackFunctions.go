package database

import (
	"fmt"
	"time"

	data "github.com/maseiler/avh-booking-system/server/data"
)

// GetFeedback returns map of feedback and its ID from database
func GetFeedback() []data.Feedback {
	feedbackMap := []data.Feedback{}
	rows, err := db.Query("SELECT * FROM feedback;")
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		feedback := data.Feedback{}
		err := rows.Scan(&feedback.ID, &feedback.TimeStamp, &feedback.Content, &feedback.Name)
		feedbackMap = append(feedbackMap, feedback)
		HandleDatabaseError(err)
	}
	err = rows.Err()
	HandleDatabaseError(err)
	return feedbackMap
}

// AddFeedback writes feddback string into database
func AddFeedback(feedback data.Feedback) {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO feedback(time_stamp, name, text) VAlUES(?, ?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	timeStampNow := time.Now().Format("2006-01-02 15:04:05")
	res, err := stmt.Exec(timeStampNow, feedback.Name, feedback.Content)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
}

// DeleteFeedback removes feedback string from database
func DeleteFeedback(id int) {
	query := fmt.Sprintf("DELETE FROM feedback WHERE id = %d;", id)
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	fmt.Println(rows)
}
