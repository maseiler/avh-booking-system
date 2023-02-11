package database

import (
	"fmt"

	data "github.com/maseiler/avh-booking-system/server/data"
)

// getItemsByQuery returns list of items as requested in string
func getItemsByQuery(query string) []data.Item {
	items := []data.Item{}
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		item := data.Item{}
		err := rows.Scan(&item.ID, &item.Name, &item.Type, &item.Size, &item.Unit, &item.Price, &item.Enabled)
		items = append(items, item)
		HandleDatabaseError(err)
	}
	err = rows.Err()
	HandleDatabaseError(err)
	// fmt.Printf("Performed item query: \"%s\"\n", query)
	return items
}

// GetAllItems returns all items from database
func GetAllItems() []data.Item {
	query := "SELECT * FROM items;"
	return getItemsByQuery(query)
}

// GetItemsOfColumnWithValue returns all items where value matches in specific column
// func GetItemsOfColumnWithValue(column string, value string) []data.Item {
// 	//TODO string conversion ?!
// 	query := ""
// 	if column == "ItemId" {
// 		intValue, _ := strconv.Atoi(value)
// 		query = fmt.Sprintf("SELECT * FROM items WHERE %s = %d;", column, intValue)
// 	} else if column == "Price" || column == "Size" {
// 		floatValue, _ := strconv.ParseFloat(value, 32)
// 		query = fmt.Sprintf("SELECT * FROM items WHERE %s = %f;", column, floatValue)
// 	} else if column == "Name" || column == "Unit" || column == "Type" {
// 		query = fmt.Sprintf("SELECT * FROM items WHERE %s = \"%s\";", column, value)
// 	} else {
// 		panic("Invalid column")
// 	}
// 	return getItemsByQuery(query)
// }

// NewItemExists returns true if item exists in database (based on Name and Size)
func NewItemExists(newItem data.Item) bool {
	queryString := fmt.Sprintf("SELECT * FROM items WHERE name = \"%s\" AND size = %.2f;", newItem.Name, newItem.Size)
	items := getItemsByQuery(queryString)
	return len(items) > 0
}

// ItemExists returns true if item with same ItemID exists in database
func ItemExists(item data.Item) bool {
	queryString := fmt.Sprintf("SELECT * FROM items WHERE id = %d;", item.ID)
	items := getItemsByQuery(queryString)
	return len(items) > 0
}

// AddItem adds a new user to database and prints info
func AddItem(newItem data.Item) {
	// todo: get info from input
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO items(name, price, size, unit, type, enabled) VAlUES(?, ?, ?, ?, ?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	//need to convert go Boolean to SQL tinyint
	var tempEnabled int
	tempEnabled = 0
	if newItem.Enabled {
		tempEnabled = 1
	}

	res, err := stmt.Exec(newItem.Name, newItem.Price, newItem.Size, newItem.Unit, newItem.Type, tempEnabled)
	HandleDatabaseError(err)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
}

// ModifyItem replaces all values of a item
func ModifyItem(item data.Item) {
	//need to convert go Boolean to SQL tinyint
	var tempEnabled int
	tempEnabled = 0
	if item.Enabled {
		tempEnabled = 1
	}
	query := fmt.Sprintf("UPDATE items SET name = \"%s\", price = \"%f\", size = \"%f\", unit = \"%s\", type = \"%s\", enabled =\"%d\" WHERE id = %d;", item.Name, item.Price, item.Size, item.Unit, item.Type, tempEnabled, item.ID)
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	fmt.Println(rows)
}

// DeleteItem deletes a item with corresponding ID from database
func DeleteItem(item data.Item) {
	query := fmt.Sprintf("DELETE FROM items WHERE id = %d;", item.ID)
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	fmt.Println(rows)
}
