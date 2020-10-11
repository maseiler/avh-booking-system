package database

import (
	"fmt"

	data "github.com/maseiler/avh-booking-system/server/data"
)

// GetItemCategories returns all item categories in database
func GetItemCategories() []data.ItemCategory {
	query := "SELECT * FROM item_categories;"
	itemCs := []data.ItemCategory{}
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		itemC := data.ItemCategory{}
		err := rows.Scan(&itemC.ID, &itemC.Name)
		itemCs = append(itemCs, itemC)
		HandleDatabaseError(err)
	}
	err = rows.Err()
	HandleDatabaseError(err)
	// fmt.Printf("Performed item query: \"%s\"\n", query)
	return itemCs
}

// AddItemCategory adds a new item category into database
func AddItemCategory(itemCategory data.ItemCategory) {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO item_categories(name) VALUES(?)")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(itemCategory.Name)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
}

// ItemCategoryExists checks if item category with the same name already exists
func ItemCategoryExists(itemCategory data.ItemCategory) bool {
	query := fmt.Sprintf("SELECT * FROM item_category WHERE name = \"%s\";", itemCategory.Name)
	categories := []data.ItemCategory{}
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		category := data.ItemCategory{}
		err := rows.Scan(category.ID, category.Name)
		categories = append(categories, category)
		HandleDatabaseError(err)
	}
	err = rows.Err()
	HandleDatabaseError(err)
	if len(categories) == 0 {
		return false
	}
	return true
}

// TODO ModifyItemCategory

// DeleteItemCategory removes item category from database
func DeleteItemCategory(itemCategory data.ItemCategory) {
	query := fmt.Sprintf("DELETE FROM item_categories WHERE id = %d;", itemCategory.ID)
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	fmt.Println(rows)
}
