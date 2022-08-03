package handler

import (
	"io"
	"net/http"
)

func GetHello(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
