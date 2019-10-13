package database

import (
	"fmt"
	"time"
)

// GetBookingStats perfroms query to return a map of the time stamp and total number of bought drinks.
func GetBookingStats(days int) map[string]int {
	m := make(map[string]int)
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	for i := 0; i < days; i++ {
		start := today.Add(time.Hour * time.Duration(-i*24))
		end := today.Add(time.Hour*time.Duration(-i*24+23) + time.Minute*time.Duration(59) + time.Second*time.Duration(59) + time.Millisecond*time.Duration(999))
		var amount int
		query := fmt.Sprintf("SELECT COALESCE(SUM(amount), 0) FROM bookings WHERE time_stamp BETWEEN \"%s\" AND \"%s\" AND item_id != 0;", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
		rows, err := db.Query(query)
		HandleDatabaseError(err)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&amount)
			HandleDatabaseError(err)
		}
		m[start.Format("2006-01-02 15:04:05")] = amount
	}
	return m
}
