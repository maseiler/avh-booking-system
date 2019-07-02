package database

import (
	"fmt"
	"strconv"
	"time"

	"../data"
)

// getBookingsFromQuery returns list of book entries as requested in string
func getBookingsFromQuery(query string) []data.BookEntry {
	bookings := []data.BookEntry{}
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		bookEntry := data.BookEntry{}
		err := rows.Scan(&bookEntry.ID, &bookEntry.TimeStamp, &bookEntry.UserID, &bookEntry.ItemID, &bookEntry.Amount, &bookEntry.TotalPrice, &bookEntry.Comment)
		bookings = append(bookings, bookEntry)
		HandleDatabaseError(err)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	// fmt.Printf("Performed booking query: \"%s\"\n", query)
	return bookings
}

func getTimestampFromQuery(query string) time.Time {
	var latestPayment time.Time
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&latestPayment)
		HandleDatabaseError(err)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	return latestPayment
}

// GetAllBookings returns all book entries in database
func GetAllBookings() []data.BookEntry {
	query := "SELECT * FROM bookings;"
	return getBookingsFromQuery(query)
}

// GetLastNBookings returns last n book entries
func GetLastNBookings(n int) []data.BookEntry {
	query := fmt.Sprintf("SELECT * FROM bookings ORDER BY id DESC LIMIT %d;", n)
	return getBookingsFromQuery(query)
}

// GetBookingsBetween returns all book entries between timespan
func GetBookingsBetween(start time.Time, end time.Time) []data.BookEntry {
	query := fmt.Sprintf("SELECT * FROM bookings WHERE time_stamp BETWEEN \"%s\" AND \"%s\";", start.Format(time.RFC3339), end.Format(time.RFC3339))
	return getBookingsFromQuery(query)
}

// GetBookingsOfColumnWithValue returns all book entries where value matches in specific column
// e.g. column="user_id" and value="12" returns all book entries of user 12
func GetBookingsOfColumnWithValue(column string, value string) []data.BookEntry {
	query := ""
	if column == "id" || column == "user_id" || column == "item_id" || column == "amount" {
		intValue, _ := strconv.Atoi(value) // TODO: error handling
		query = fmt.Sprintf("SELECT * FROM bookings Where %s = %d;", column, intValue)
	} else if column == "time_stamp" {
		//TODO func CheckTimePattern
		query = fmt.Sprintf("SELECT * FROM bookings Where %s = \"%s\";", column, value)
		//TODO client side: convert to format
	} else if column == "total_price" {
		floatValue, _ := strconv.ParseFloat(value, 32)
		query = fmt.Sprintf("SELECT * FROM bookings Where %s = %f;", column, floatValue)
	} else if column == "comment" {
		query = fmt.Sprintf("SELECT * FROM bookings Where %s = \"%s\";", column, value)
	} else {
		panic("Invalid column name")
	}
	return getBookingsFromQuery(query)
}

// GetUserDebts returns list of book entries which have not yet payed by user
func GetUserDebts(user data.User) data.Debts {
	lastPayDayQuery := fmt.Sprintf("SELECT time_stamp FROM bookings WHERE user_id = %d AND total_price <= 0 ORDER BY time_stamp DESC LIMIT 1;", user.ID)
	lastPayDay := getTimestampFromQuery(lastPayDayQuery)
	debtsQuery := fmt.Sprintf("SELECT * FROM bookings WHERE user_id = %d AND time_stamp > \"%s\";", user.ID, lastPayDay.Format(time.RFC3339))
	unpaid := getBookingsFromQuery(debtsQuery)
	debts := data.Debts{LastPayment: lastPayDay, Debts: unpaid}
	return debts
}

// Checkout adds a Cart to bookings and updates favorite items in database.
func Checkout(cart data.Cart) bool {
	if cart.User.Balance >= float32(cart.User.MaxDebt) {
		return false
	}
	numItems := len(cart.CartItems)
	for i := 0; i < numItems; i++ {
		tx, err := db.Begin()
		HandleDatabaseError(err)
		stmt, err := tx.Prepare("INSERT INTO bookings(time_stamp, user_id, item_id, amount, total_price, comment) VAlUES(?, ?, ?, ?, ?, ?)")
		HandleTxError(tx, err)
		defer stmt.Close()
		timeStamp := time.Now().Format(time.RFC3339)
		totalPrice := float32(cart.CartItems[i].Amount) * cart.CartItems[i].Item.Price
		comment := fmt.Sprintf("Part %d/%d", i+1, numItems)
		res, err := stmt.Exec(timeStamp, cart.User.ID, cart.CartItems[i].Item.ID, cart.CartItems[i].Amount, totalPrice, comment)
		TxRowsAffected(res, tx)
		err = tx.Commit()
		HandleDatabaseError(err)
		stmt.Close()

		user := GetUsersOfColumnWithValue("id", strconv.Itoa(cart.User.ID))[0]
		user.Balance += totalPrice
		ModifyUser(user)
	}
	UpdateFavoriteItems(cart)
	return true
}

// Pay creates a book entry with inverted balance and sets user balance to 0.
func Pay(user data.User) bool {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO bookings(time_stamp, user_id, item_id, amount, total_price, comment) VAlUES(?, ?, ?, ?, ?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	timeStamp := time.Now().Format(time.RFC3339)
	totalPrice := -float32(user.Balance)
	comment := "Payment"
	res, err := stmt.Exec(timeStamp, user.ID, 1, 1, totalPrice, comment)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	stmt.Close()

	query := fmt.Sprintf("UPDATE users SET balance = 0 WHERE id = %d;", user.ID)
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	fmt.Println(rows)
	if err == nil {
		return true
	}
	return false
}

// DeleteBookEntry deletes an entry from database.
func DeleteBookEntry(entry data.BookEntry) bool {
	user := GetUsersOfColumnWithValue("id", strconv.Itoa(entry.UserID))[0]
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("DELETE FROM bookings WHERE id = ?")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(entry.ID)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	stmt.Close()
	if err == nil {
		if entry.TotalPrice > 0 {
			user.Balance -= entry.TotalPrice
		} else if entry.TotalPrice < 0 {
			user.Balance -= entry.TotalPrice
		}
		ModifyUser(user)
		return true
	}
	return false
}
