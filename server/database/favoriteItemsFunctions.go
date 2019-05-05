package database

import (
	"fmt"

	data "../data"
)

func entryExists(userID int, itemID int) bool {
	query := fmt.Sprintf("SELECT EXISTS( SELECT 1 FROM favoriteItems WHERE UserId = %d AND ItemId = %d);", userID, itemID)
	var exists bool
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&exists)
		HandleDatabaseError(err)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	return exists
}

func addEntry(userID int, ItemID int, amount int) error {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	stmt, err := tx.Prepare("INSERT INTO favoriteItems (UserId, ItemId, Count) VALUES (?, ?, ?);")
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec(userID, ItemID, amount)
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	stmt.Close()
	return err
}

func incrementCount(userID int, ItemID int, amount int) error {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	query := fmt.Sprintf("UPDATE favoriteItems SET Count = Count+%d WHERE UserId = %d AND ItemId = %d;", amount, userID, ItemID)
	stmt, err := tx.Prepare(query)
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec()
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	stmt.Close()
	return err
}

// UpdateFavoriteItems creates a new entry in database if item and user doesn't exists or increments count otherwise.
func UpdateFavoriteItems(cart data.Cart) bool {
	numItems := len(cart.CartItems)
	var err error
	for i := 0; i < numItems; i++ {
		if entryExists(cart.User.UserID, cart.CartItems[i].Item.ItemID) {
			err = incrementCount(cart.User.UserID, cart.CartItems[i].Item.ItemID, cart.CartItems[i].Amount)
		} else {
			err = addEntry(cart.User.UserID, cart.CartItems[i].Item.ItemID, cart.CartItems[i].Amount)
		}
		if err != nil {
			return false
		}
	}
	return true
}

// GetFavoriteItemIDs returns list of favorite item IDs from databse.
func GetFavoriteItemIDs(userID int) []int {
	query := fmt.Sprintf("SELECT ItemId FROM favoriteItems WHERE UserId = %d ORDER BY Count DESC LIMIT 5;", userID)
	var favoriteItemIDs []int
	rows, err := db.Query(query)
	HandleDatabaseError(err)
	defer rows.Close()
	for rows.Next() {
		var itemID int
		err := rows.Scan(&itemID)
		HandleDatabaseError(err)
		favoriteItemIDs = append(favoriteItemIDs, itemID)
	}
	defer rows.Close()
	err = rows.Err()
	HandleDatabaseError(err)
	return favoriteItemIDs
}

// DeleteUserFromFavoriteItems deletes all entries in favoriteItems of user.
func DeleteUserFromFavoriteItems(user data.User) bool {
	tx, err := db.Begin()
	HandleDatabaseError(err)
	query := fmt.Sprintf("DELETE FROM favoriteItems WHERE UserId = %d;", user.UserID)
	stmt, err := tx.Prepare(query)
	HandleTxError(tx, err)
	defer stmt.Close()
	res, err := stmt.Exec()
	TxRowsAffected(res, tx)
	err = tx.Commit()
	HandleDatabaseError(err)
	stmt.Close()
	if err != nil {
		return true
	}
	return false
}

//TODO create: PerformTransaction(query string)
