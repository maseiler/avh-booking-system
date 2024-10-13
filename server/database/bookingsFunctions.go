package database

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/maseiler/avh-booking-system/server/data"
)

// getBookEntriesFromQuery returns list of book entries as requested in string
func getBookEntriesFromQuery(query string) []data.BookEntry {
	bookEntries := []data.BookEntry{}
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		bookEntry := data.BookEntry{}
		err := rows.Scan(&bookEntry.ID, &bookEntry.TimeStamp, &bookEntry.UserID, &bookEntry.ItemID, &bookEntry.Amount, &bookEntry.TotalPrice, &bookEntry.Comment, &bookEntry.PaymentMethod)
		bookEntries = append(bookEntries, bookEntry)
		HandleDatabaseError(err)
	}
	err = rows.Err()
	HandleDatabaseError(err)
	// fmt.Printf("Performed booking query: \"%s\"\n", query)
	return bookEntries
}

// GetLastNBookEntries returns the last n book entries
func GetLastNBookEntries(n int) []data.BookEntry {
	query := fmt.Sprintf("SELECT * FROM bookings ORDER BY id DESC LIMIT %d;", n)
	return getBookEntriesFromQuery(query)
}

// GetBookEntriesBetween returns all book entries within timespan
func GetBookEntriesBetween(start time.Time, end time.Time) []data.BookEntry {
	query := fmt.Sprintf("SELECT * FROM bookings WHERE time_stamp BETWEEN \"%s\" AND \"%s\" ORDER BY id DESC;", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
	if (start == time.Time{}) || (end == time.Time{}) {
		query = "SELECT * FROM bookings ORDER BY id DESC;"
	}
	return getBookEntriesFromQuery(query)
}

// GetBookEntriesOfUserBetween returns all book entries of specified user within timespan
func GetBookEntriesOfUserBetween(user data.User, start time.Time, end time.Time) []data.BookEntry {
	query := fmt.Sprintf("SELECT * FROM bookings WHERE user_id = %d AND time_stamp BETWEEN \"%s\" AND \"%s\" ORDER BY id DESC;", user.ID, start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
	if (start == time.Time{}) || (end == time.Time{}) {
		query = fmt.Sprintf("SELECT * FROM bookings WHERE user_id = %d ORDER BY id DESC;", user.ID)
	}
	return getBookEntriesFromQuery(query)
}

// GetBookEntriesOfItemBetween returns all book entries of specified item within timespan
func GetBookEntriesOfItemBetween(item data.Item, start time.Time, end time.Time) []data.BookEntry {
	query := fmt.Sprintf("SELECT * FROM bookings WHERE item_id = %d AND time_stamp BETWEEN \"%s\" AND \"%s\" ORDER BY id DESC;", item.ID, start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
	if (start == time.Time{}) || (end == time.Time{}) {
		query = fmt.Sprintf("SELECT * FROM bookings WHERE item_id = %d ORDER BY id DESC;", item.ID)
	}
	return getBookEntriesFromQuery(query)
}

// GetPaymentsOfUser returns all payment of specified user within timespan
func GetPaymentsOfUser(user data.User, start time.Time, end time.Time) []data.BookEntry {
	emptyUser := data.User{ID: 0, BierName: "", FirstName: "", LastName: "", BoatName: "", Status: "", Email: "", Phone: "", Balance: 0, MaxDebt: 0}
	var query string
	if user == emptyUser {
		query = fmt.Sprintf("SELECT * from bookings WHERE total_price <= 0 AND time_stamp BETWEEN \"%s\" AND \"%s\" ORDER BY id DESC;", start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
		if (start == time.Time{}) || (end == time.Time{}) {
			query = "SELECT * from bookings WHERE total_price <= 0 ORDER BY id DESC;"
		}
	} else {
		query = fmt.Sprintf("SELECT * from bookings WHERE user_id = %d AND total_price <= 0 AND time_stamp BETWEEN \"%s\" AND \"%s\" ORDER BY id DESC;", user.ID, start.Format("2006-01-02 15:04:05"), end.Format("2006-01-02 15:04:05"))
		if (start == time.Time{}) || (end == time.Time{}) {
			query = fmt.Sprintf("SELECT * from bookings WHERE user_id = %d AND total_price <= 0 ORDER BY id DESC;", user.ID)
		}
	}
	return getBookEntriesFromQuery(query)
}

// GetUserDebts returns list of book entries which have not yet paid by user
func GetUserDebts(user data.User) data.Debts {
	lastPayDayQuery := fmt.Sprintf("SELECT time_stamp FROM bookings WHERE user_id = %d AND comment = 'Payment Full' ORDER BY time_stamp DESC LIMIT 1;", user.ID)
	lastPayDay := getTimestampFromQuery(lastPayDayQuery)
	debtsQuery := fmt.Sprintf("SELECT * FROM bookings WHERE user_id = %d AND time_stamp > \"%s\" ORDER BY id DESC;", user.ID, lastPayDay.Format("2006-01-02 15:04:05"))
	unpaid := getBookEntriesFromQuery(debtsQuery)
	debts := data.Debts{LastPayment: lastPayDay, Debts: unpaid}
	return debts
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
	err = rows.Err()
	HandleDatabaseError(err)
	return latestPayment
}

// Checkout adds a Cart to bookings and updates favorite items in database.
func Checkout(cart data.Cart) bool {
	if cart.User.Balance >= float32(cart.User.MaxDebt) {
		return false
	}
	numItems := len(cart.CartItems)
	for i := 0; i < numItems; i++ {
		tx, err := db.Begin()
		if err != nil {
			log.Println(err)
			return false
		}
		HandleDatabaseError(err)
		stmt, err := tx.Prepare("INSERT INTO bookings(time_stamp, user_id, item_id, amount, total_price, comment, payment_method) VAlUES(?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Println(err)
			return false
		}
		HandleTxError(tx, err)
		defer stmt.Close()
		timeStamp := time.Now().Format("2006-01-02 15:04:05")
		totalPrice := float32(cart.CartItems[i].Amount) * cart.CartItems[i].Item.Price
		comment := fmt.Sprintf("Part %d/%d", i+1, numItems)
		paymentMethod := ""
		res, err := stmt.Exec(timeStamp, cart.User.ID, cart.CartItems[i].Item.ID, cart.CartItems[i].Amount, totalPrice, comment, paymentMethod)
		if err != nil {
			log.Println(err)
			return false
		}
		TxRowsAffected(res, tx)
		err = tx.Commit()
		if err != nil {
			log.Println(err)
			return false
		}
		HandleDatabaseError(err)

		user := GetUsersOfColumnWithValue("id", strconv.Itoa(cart.User.ID))[0]
		user.Balance += totalPrice
		ModifyUser(user)
	}
	UpdateFavoriteItems(cart)
	return true
}

// Pay creates a book entry with inverted balance and sets user balance to 0.
func Pay(payment data.Payment) bool {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO bookings(time_stamp, user_id, item_id, amount, total_price, comment, payment_method) VAlUES(?, ?, ?, ?, ?, ?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	totalPrice := -float32(payment.Balance)
	comment := "Payment Part"
	if payment.Balance == payment.User.Balance {
		comment = "Payment Full"
	}
	res, err := stmt.Exec(timeStamp, payment.User.ID, 0, 1, totalPrice, comment, payment.PaymentMethod)
	HandleDatabaseError(err)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)

	// Set new User.Balance based on the payment amount.
	newBalance := payment.User.Balance - payment.Balance
	query := fmt.Sprintf("UPDATE users SET balance = %.2f WHERE id = %d;", newBalance, payment.User.ID)
	_, err = db.Query(query)
	HandleDatabaseError(err)
	return err == nil
}

// DeleteBookEntry deletes an entry from database.
func DeleteBookEntry(entry data.BookEntry) bool {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("DELETE FROM bookings WHERE id = ?")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(entry.ID)
	HandleDatabaseError(err)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	return err == nil
}

// UndoBookEntry creates a new book entry with inversed balance and adjusts the user's balance accordingly.
func UndoBookEntry(entry data.BookEntry) bool {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO bookings(time_stamp, user_id, item_id, amount, total_price, comment, payment_method) VAlUES(?, ?, ?, ?, ?, ?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	totalPrice := -float32(entry.TotalPrice)
	comment := fmt.Sprintf("Undo book entry %d", entry.ID)
	res, err := stmt.Exec(timeStamp, entry.UserID, 0, 1, totalPrice, comment, "")
	HandleDatabaseError(err)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)

	user := GetUsersOfColumnWithValue("id", strconv.Itoa(entry.UserID))[0]
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
