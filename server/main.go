package main

import (
	"net/http"

	db "./database"
	handler "./handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db.ConnectDatabase()

	// api
	r := mux.NewRouter()
	r.HandleFunc("/addUser", handler.AddUser)
	r.HandleFunc("/users", handler.GetUsers)
	r.HandleFunc("/newUser", handler.AddUser)
	r.HandleFunc("/items", handler.GetItems)
	r.HandleFunc("/checkout", handler.Checkout)
	// frontend
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./../client/dist")))
	http.ListenAndServe(":8080", r)

}
