package handler

import (
	"fmt"
	"net/http"

	dbP "../database"
)

// UpdateFavoriteItems forwards API calll to create a new entry in database if item and user doesn't exists or increments count otherwise.
func UpdateFavoriteItems(w http.ResponseWriter, r *http.Request) {
	cart := UnmarshalCart(r.Body)
	validation := ""
	if dbP.UpdateFavoriteItems(cart) {
		validation = "ok"
		w.WriteHeader(http.StatusOK)
	} else {
		validation = fmt.Sprintf("Couldn't update favorite items for user %d.", cart.User.UserID)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprint(w, validation)
}

// GetFavoriteItemIDs forwards API call to get the list of favorite item IDs of user
func GetFavoriteItemIDs(w http.ResponseWriter, r *http.Request) {
	user := UnmarshalUser(r.Body)
	favoriteItems := dbP.GetFavoriteItemIDs(user.UserID)
	response := marshalToJSON(favoriteItems, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// DeleteUserFromFavoriteItems forwards API call to delete all entries in favoriteItems of user.
func DeleteUserFromFavoriteItems(w http.ResponseWriter, r *http.Request) {
	user := UnmarshalUser(r.Body)
	validation := ""
	if dbP.DeleteUserFromFavoriteItems(user) {
		validation = "ok"
		w.WriteHeader(http.StatusOK)
	} else {
		validation = fmt.Sprintf("Couldn't delete user %d from favoriteItems.", user.UserID)
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Fprint(w, validation)
}
