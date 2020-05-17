package database

import (
	"fmt"
	"strconv"
	"time"

	"github.com/maseiler/avh-booking-system/server/data"
)

// GetBookingStats performs query to return a map of time stamps, total buying and buying per item.
func GetBookingStats(days int) map[string][]string {
	// get list of item IDs
	var itemIDs []int
	query := fmt.Sprintf("SELECT id FROM items;")
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		HandleDatabaseError(err)
		itemIDs = append(itemIDs, id)
	}

	m := make(map[string][]string)
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	for i := 0; i < days; i++ {
		start := today.Add(time.Hour * time.Duration(-i*24))
		end := today.Add(time.Hour*time.Duration(-i*24+23) + time.Minute*time.Duration(59) + time.Second*time.Duration(59) + time.Millisecond*time.Duration(999))
		var amount int

		// get total bookings
		query := fmt.Sprintf("SELECT COALESCE(SUM(amount), 0) FROM bookings WHERE time_stamp BETWEEN \"%s\" AND \"%s\" AND item_id != 0;", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
		rows, err := db.Query(query)
		HandleDatabaseError(err)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&amount)
			HandleDatabaseError(err)
		}

		m["timeStamp"] = append(m["timeStamp"], start.Format("2006-01-02 15:04:05"))
		m["total"] = append(m["total"], strconv.Itoa(amount))

		// get bookings of item
		for _, id := range itemIDs {
			query := fmt.Sprintf("SELECT COALESCE(SUM(amount), 0) FROM bookings WHERE time_stamp BETWEEN \"%s\" AND \"%s\" AND item_id = %d;", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"), id)
			rows, err := db.Query(query)
			HandleDatabaseError(err)
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&amount)
				HandleDatabaseError(err)
			}
			m[strconv.Itoa(id)] = append(m[strconv.Itoa(id)], strconv.Itoa(amount))
		}
	}
	return m
}

// GetFavoriteItemsStats perform query to return a list ItemStats (if count greater than 0)
func GetFavoriteItemsStats() []data.ItemStat {
	m := make(map[string]int)
	// get maxID to determine end of for loop
	var maxID int
	query := fmt.Sprintf("SELECT MAX(id) FROM items;")
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&maxID)
		HandleDatabaseError(err)
	}

	// get count of item
	for i := 1; i <= maxID; i++ {
		var amount int
		query = fmt.Sprintf("SELECT COALESCE(SUM(count), 0) FROM favorite_items WHERE item_id = %d;", i)
		rows, err := db.Query(query)
		HandleDatabaseError(err)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&amount)
			HandleDatabaseError(err)
		}
		m[strconv.Itoa(i)] = amount
	}

	// delete all with count 0
	for k := range m {
		if m[k] == 0 {
			delete(m, k)
		}
	}

	stats := []data.ItemStat{}
	// get additional item info
	for k := range m {
		itemStat := data.ItemStat{}
		id, err := strconv.Atoi(k)
		query = fmt.Sprintf("SELECT name, size, unit, type FROM items WHERE id = %d;", id)
		rows, err := db.Query(query)
		HandleDatabaseError(err)
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&itemStat.Name, &itemStat.Size, &itemStat.Unit, &itemStat.Type)
			HandleDatabaseError(err)
		}
		itemStat.Count = m[k]
		stats = append(stats, itemStat)
	}

	return stats
}
