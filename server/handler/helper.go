package handler

import (
	"bytes"
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

// UnmarshalString unmarshals JSON object to string
func UnmarshalString(body io.ReadCloser) string {
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(body)
	str := buf.String()

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return str
}

// UnmarshalStringArray unmarshals JSON object to string array
func UnmarshalStringArray(body io.ReadCloser) []string {
	var array []string
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(body)
	err = json.Unmarshal(buf.Bytes(), &array)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return array
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

// UnmarshalBookEntry unmarshals JSON object to BookEntry
func UnmarshalBookEntry(body io.ReadCloser) dataP.BookEntry {
	decoder := json.NewDecoder(body)

	var entry dataP.BookEntry
	err := decoder.Decode(&entry)

	if err != nil {
		panic(err)
	}
	return entry
}

// UnmarshalFeedback unmarshals JSON object to Feedback
func UnmarshalFeedback(body io.ReadCloser) dataP.Feedback {
	decoder := json.NewDecoder(body)

	var feedback dataP.Feedback
	err := decoder.Decode(&feedback)

	if err != nil {
		panic(err)
	}
	return feedback
}

func handleHTTPError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
