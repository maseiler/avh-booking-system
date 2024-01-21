package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	dataP "github.com/maseiler/avh-booking-system/server/data"
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

// UnmarshalItemFromTo unmarshals JSON object to ItemFromTo struct
func UnmarshalItemFromTo(body io.ReadCloser) dataP.ItemFromTo {
	var itemFromTo dataP.ItemFromTo
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(body)
	err = json.Unmarshal(buf.Bytes(), &itemFromTo)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return itemFromTo
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

// UnmarshalUserFromTo unmarshals JSON object to UserFromTo struct
func UnmarshalUserFromTo(body io.ReadCloser) dataP.UserFromTo {
	var userFromTo dataP.UserFromTo
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(body)
	err = json.Unmarshal(buf.Bytes(), &userFromTo)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return userFromTo
}

// UnmarshalUserDouble unmarshals JSON object to UserInt struct
func UnmarshalUserDouble(body io.ReadCloser) dataP.UserDouble {
	var userDouble dataP.UserDouble
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(body)
	err = json.Unmarshal(buf.Bytes(), &userDouble)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return userDouble
}

// UnmarshalPayment unmarshals JSON object to Payment struct
func UnmarshalPayment(body io.ReadCloser) dataP.Payment {
	var payment dataP.Payment
	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(body)
	err = json.Unmarshal(buf.Bytes(), &payment)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return payment
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

// UnmarshalFeedback unmarshals JSON object to Feedback
func UnmarshalSetting(body io.ReadCloser) dataP.Setting {
	decoder := json.NewDecoder(body)

	var setting dataP.Setting
	err := decoder.Decode(&setting)

	if err != nil {
		panic(err)
	}
	return setting
}

// ConvertStringToTime converts string of format YYY-MM-DD to time
func ConvertStringToTime(str string) time.Time {
	layout := "2006-01-02"
	t, err := time.Parse(layout, str)

	if err != nil {
		return time.Time{}
	}

	return t
}

func handleHTTPError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
