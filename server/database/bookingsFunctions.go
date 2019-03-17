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
		err := rows.Scan(&bookEntry.BookEntryID, &bookEntry.TimeStamp, &bookEntry.UserID, &bookEntry.ItemID, &bookEntry.Amount, &bookEntry.TotalPrice, &bookEntry.Comment)
		bookings = append(bookings, bookEntry)
		HandleDatabaseError(err)
		// info := fmt.Sprintf("BookEntryId: %d\nTimeStamp: %s\nUserID: %d\nItemID: %d\nAmoint: %d\nTotalPrice: %f\nComment: %s\n", bookEntry.BookEntryID, bookEntry.TimeStamp, bookEntry.UserID, bookEntry.ItemID, bookEntry.Amount, bookEntry.TotalPrice, bookEntry.Comment)
		// fmt.Println(info)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	fmt.Printf("Performed booking query: \"%s\"\n", query)
	return bookings
}

// GetAllBookings returns all book entries in database
func GetAllBookings() []data.BookEntry {
	query := "SELECT * FROM bookings;"
	return getBookingsFromQuery(query)
}

// GetLastNBookings returns last n book entries
func GetLastNBookings(n int) []data.BookEntry {
	query := fmt.Sprintf("SELECT * FROM bookings ORDER BY BookEntryId DESC LIMIT %d;", n)
	return getBookingsFromQuery(query)
}

// GetBookingsBetween returns all book entries between timespan
func GetBookingsBetween(start time.Time, end time.Time) []data.BookEntry {
	query := fmt.Sprintf("SELECT * FROM bookings WHERE TimeStamp BETWEEN \"%s\" AND \"%s\";", start.Format(time.RFC3339), end.Format(time.RFC3339))
	return getBookingsFromQuery(query)
}

// GetBookingsOfColumnWithValue returns all book entries where value matches in specific column
// e.g. column="UserID" and value="12" returns all book entries of user 12
func GetBookingsOfColumnWithValue(column string, value string) []data.BookEntry {
	query := ""
	if column == "BookEntryId" || column == "UserId" || column == "ItemId" || column == "Amount" {
		intValue, _ := strconv.Atoi(value) // TODO: error handling
		query = fmt.Sprintf("SELECT * FROM bookings Where %s = %d;", column, intValue)
	} else if column == "TimeStamp" {
		//TODO func CheckTimePattern
		query = fmt.Sprintf("SELECT * FROM bookings Where %s = \"%s\";", column, value)
		//TODO client side: convert to format
	} else if column == "TotalPrice" {
		floatValue, _ := strconv.ParseFloat(value, 32)
		query = fmt.Sprintf("SELECT * FROM bookings Where %s = %f;", column, floatValue)
	} else if column == "Comment" {
		query = fmt.Sprintf("SELECT * FROM bookings Where %s = \"%s\";", column, value)
	} else {
		panic("Invalid column name")
	}
	return getBookingsFromQuery(query)
}
