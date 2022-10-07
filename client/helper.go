package main

import (
	"encoding/json"
	"io"
)

func UnmarshalUser(body io.ReadCloser) User {
	decoder := json.NewDecoder(body)

	var user User
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}
	return user
}

func UnmarshalUserList(body io.ReadCloser) []User {
	decoder := json.NewDecoder(body)

	var users []User
	err := decoder.Decode(&users)

	if err != nil {
		panic(err)
	}
	return users
}

func HandleMarshalingError(err error) {
	if err != nil {
		panic(err)
	}
}
