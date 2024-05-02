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
	query := "SELECT id FROM items;"
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

		// get total sum of book entries
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

		// get book entries of item
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
	query := "SELECT MAX(id) FROM items;"
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
		id, _ := strconv.Atoi(k)
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

// Get All Debts and Credits based on the Categorys
func GetAllCategoryDebts(until string) map[string]float32 {
	var (
		status string
		debt   float32
	)
	queryString := fmt.Sprintf("SELECT users.status, SUM(user_debts) AS debts FROM (SELECT user_id, SUM(total_price) AS user_debts FROM bookings WHERE time_stamp < '%s' GROUP BY user_id) subq JOIN users ON subq.user_id = users.id GROUP BY status;", until)
	rows, err := db.Query(queryString)
	HandleDatabaseError(err)
	result := make(map[string]float32)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&status, &debt)
		HandleDatabaseError(err)
		result[status] = debt
	}
	err = rows.Err()
	HandleDatabaseError(err)
	return result
}

// Get amount of sold items in the last week
// Ich krig den Rückgabewert nicht ordentlich hin... Brauche hilfe mit go.
func GetLasWeekSoldItems() {
	// var (
	// 	Ware    string
	// 	Groeße  string
	// 	Einheit string
	// 	Anzahl  string
	// )
	// // ToDo: Reduziere die Abfrage auf die letzte Woche - aktuell wird noch alles abgefragt
	// queryString := "SELECT items.name as 'Ware', items.size as 'Groeße', items.unit as 'Einheit' , SUM(amount) as 'Anzahl' FROM bookings LEFT JOIN items on item_id=items.id GROUP BY item_id;"
	// rows, err := db.Query(queryString)
	// HandleDatabaseError(err)
	// var result [][]
	// defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan(&Ware, &Groeße, &Einheit, &Anzahl)
	// 	HandleDatabaseError(err)
	// 	row := [4]string{Ware, Groeße, Einheit, Anzahl}
	// 	result = append(result, row)
	// }
	// err = rows.Err()
	// HandleDatabaseError(err)
	// return result
}
