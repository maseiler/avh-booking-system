package database

import (
	"fmt"

	data "../data"
)

// GetFeedback returns map  of feedback and its ID from database
func GetFeedback() []data.Feedback {
	feedbackMap := []data.Feedback{}
	rows, err := db.Query("SELECT * FROM feedback;")
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		feedback := data.Feedback{}
		err := rows.Scan(&feedback.ID, &feedback.Content)
		feedbackMap = append(feedbackMap, feedback)
		HandleDatabaseError(err)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	return feedbackMap
}

// AddFeedback writes feddback string into database
func AddFeedback(text string) {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO feedback(Text) VAlUES(?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(text)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	stmt.Close()
}

// DeleteFeedback removes feedback string from database
func DeleteFeedback(id int) {
	query := fmt.Sprintf("DELETE FROM feedback WHERE ID = %d;", id)
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	fmt.Println(rows)
}
