package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	r.HandleFunc("/getUserDebts", handler.GetUserDebts)
	r.HandleFunc("/checkout", handler.Checkout)
	r.HandleFunc("/pay", handler.Pay)
	r.HandleFunc("/deleteBookEntry", handler.DeleteBookEntry)
	r.HandleFunc("/updateFavoriteItems", handler.UpdateFavoriteItems)
	r.HandleFunc("/getFavoriteItemIDs", handler.GetFavoriteItemIDs)
	r.HandleFunc("/deleteUserFromFavoriteItems", handler.DeleteUserFromFavoriteItems)
	r.HandleFunc("/getFeedback", handler.GetFeedback)
	r.HandleFunc("/addFeedback", handler.AddFeedback)
	r.HandleFunc("/deleteFeedback", handler.DeleteFeedback)
	r.HandleFunc("/login", handler.Login)
	r.HandleFunc("/changeAdminPassword", handler.ChangeAdminPassword)

	serveIndexHTML := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		http.ServeFile(w, r, filepath.Join(os.Getenv("AVHBS_FRONTEND_PATH"), "index.html"))
	}
	r.PathPrefix("/").Handler(handler.CustomFileServer(http.Dir(os.Getenv("AVHBS_FRONTEND_PATH")), serveIndexHTML))

	server := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServeTLS(os.Getenv("AVHBS_SSL_CRT_FILE_PATH"), os.Getenv("AVHBS_SSL_KEY_FILE_PATH")))
}
