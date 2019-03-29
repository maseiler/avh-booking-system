package main

import (
	"log"
	"net/http"
	"time"

	db "./database"
	handler "./handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db.ConnectDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/getUsers", handler.GetUsers)
	r.HandleFunc("/addUser", handler.AddUser)
	r.HandleFunc("/modifyUser", handler.ModifyUser)
	r.HandleFunc("/deleteUser", handler.DeleteUser)
	r.HandleFunc("/getUnreservedItems", handler.GetUnreservedItems)
	r.HandleFunc("/getReservedItems", handler.GetReservedItems)
	r.HandleFunc("/addItem", handler.AddItem)
	r.HandleFunc("/modifyItem", handler.ModifyItem)
	r.HandleFunc("/deleteItem", handler.DeleteItem)
	r.HandleFunc("/getLastNBookings", handler.GetLastNBookings)
	r.HandleFunc("/checkout", handler.Checkout)
	serveIndexHTML := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./../client/dist/index.html")
	}
	r.PathPrefix("/").Handler(handler.CustomFileServer(http.Dir("./../client/dist"), serveIndexHTML))

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(server.ListenAndServe())
}
