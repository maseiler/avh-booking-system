package main

import (
	"log"
	"net/http"
	"os"
	"time"

	db "github.com/maseiler/avh-booking-system/server/database"
	handler "github.com/maseiler/avh-booking-system/server/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db.CreateDatabase()

	r := mux.NewRouter()
	// user routes
	r.HandleFunc("/addNewUser", handler.AddNewUser)
	r.HandleFunc("/getAllUsers", handler.GetAllUsers)

	server := &http.Server{
		Addr:           ":8081",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if len(os.Getenv("AVHBS_SSL_CRT_FILE_PATH")) > 0 {
		log.Fatal(server.ListenAndServeTLS(os.Getenv("AVHBS_SSL_CRT_FILE_PATH"), os.Getenv("AVHBS_SSL_KEY_FILE_PATH")))
	} else {
		log.Fatal(server.ListenAndServe())
	}
}
