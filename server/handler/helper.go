package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	dataP "../data"
)

func marshalToJSON(object interface{}, w http.ResponseWriter) (response []byte) {
	response, err := json.Marshal(object)
	handleHTTPError(err, w)
	return response
}

// UnmarshalInt unmarshals JSON object to int
func UnmarshalInt(body io.ReadCloser) int {
	decoder := json.NewDecoder(body)
	var n int
	err := decoder.Decode(&n)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return n
}

// UnmarshalItem unmarshals JSON object to Item
func UnmarshalItem(body io.ReadCloser) dataP.Item {
	decoder := json.NewDecoder(body)

	var item dataP.Item
	err := decoder.Decode(&item)

	if err != nil {
		panic(err)
	}
	return item
}

// UnmarshalUser unmarshals JSON object to User
func UnmarshalUser(body io.ReadCloser) dataP.User {
	decoder := json.NewDecoder(body)

	var user dataP.User
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}
	return user
}

// UnmarshalCart unmarshals JSON object to Cart
func UnmarshalCart(body io.ReadCloser) dataP.Cart {
	decoder := json.NewDecoder(body)

	var cart dataP.Cart
	err := decoder.Decode(&cart)

	if err != nil {
		panic(err)
	}
	return cart
}
func handleHTTPError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
