package database

import (
	"fmt"

	data "github.com/maseiler/avh-booking-system/server/data"
)

// GetItemCategoryMaps returns all ItemCategoryMaps in database
func GetItemCategoryMaps() []data.ItemCategoryMap {
	query := "SELECT * FROM item_category_map;"
	itemCategoryMaps := []data.ItemCategoryMap{}
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		icMap := data.ItemCategoryMap{}
		err := rows.Scan(&icMap.ID, &icMap.CategoryID, &icMap.ItemID)
		itemCategoryMaps = append(itemCategoryMaps, icMap)
		HandleDatabaseError(err)
	}
	err = rows.Err()
	HandleDatabaseError(err)
	// fmt.Printf("Performed item query: \"%s\"\n", query)
	return itemCategoryMaps
}

// AddItemCategoryMap adds a new ItemCategoryMap into database
func AddItemCategoryMap(icMap data.ItemCategoryMap) {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO item_category_map(category_id, item_id) VALUES(?, ?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(icMap.CategoryID, icMap.ItemID)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
}

// ItemCategoryMapExists checks if item category map already exists
func ItemCategoryMapExists(icMap data.ItemCategoryMap) bool {
	query := fmt.Sprintf("SELECT * FROM item_category_map WHERE item_id = %d AND category_id = %d;", icMap.ItemID, icMap.CategoryID)
	icMaps := []data.ItemCategoryMap{}
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		icMap := data.ItemCategoryMap{}
		err := rows.Scan(icMap.ID, icMap.CategoryID, icMap.ItemID)
		icMaps = append(icMaps, icMap)
		HandleDatabaseError(err)
	}
	err = rows.Err()
	HandleDatabaseError(err)
	if len(icMaps) == 0 {
		return false
	}
	return true
}

// TODO ModifyItemCategoryMap

// DeleteItemCategoryMap removes ItemCategoryMap from database
func DeleteItemCategoryMap(icMap data.ItemCategoryMap) {
	query := fmt.Sprintf("DELETE FROM item_category_map WHERE id = %d;", icMap.ID)
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	fmt.Println(rows)
}
