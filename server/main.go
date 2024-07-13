package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	db "github.com/maseiler/avh-booking-system/server/database"
	handler "github.com/maseiler/avh-booking-system/server/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db.CreateDatabase()

	r := mux.NewRouter()
	r.HandleFunc("/getUsers", handler.GetUsers)
	r.HandleFunc("/addUser", handler.AddUser)
	r.HandleFunc("/modifyUser", handler.ModifyUser)
	r.HandleFunc("/deleteUser", handler.DeleteUser)
	r.HandleFunc("/getAllItems", handler.GetAllItems)
	r.HandleFunc("/addItem", handler.AddItem)
	r.HandleFunc("/modifyItem", handler.ModifyItem)
	r.HandleFunc("/deleteItem", handler.DeleteItem)
	r.HandleFunc("/getLastNBookEntries", handler.GetLastNBookEntries)
	r.HandleFunc("/getBookEntriesFromUserBetween", handler.GetBookEntriesFromUserBetween)
	r.HandleFunc("/getBookEntriesFromItemBetween", handler.GetBookEntriesFromItemBetween)
	r.HandleFunc("/getPaymentsOfUser", handler.GetPaymentsOfUser)
	r.HandleFunc("/getUserDebts", handler.GetUserDebts)
	r.HandleFunc("/checkout", handler.Checkout)
	r.HandleFunc("/pay", handler.Pay)
	r.HandleFunc("/deleteBookEntry", handler.DeleteBookEntry)
	r.HandleFunc("/undoBookEntry", handler.UndoBookEntry)
	r.HandleFunc("/updateFavoriteItems", handler.UpdateFavoriteItems)
	r.HandleFunc("/getFavoriteItemIDs", handler.GetFavoriteItemIDs)
	r.HandleFunc("/deleteUserFromFavoriteItems", handler.DeleteUserFromFavoriteItems)
	r.HandleFunc("/getFeedback", handler.GetFeedback)
	r.HandleFunc("/addFeedback", handler.AddFeedback)
	r.HandleFunc("/deleteFeedback", handler.DeleteFeedback)
	r.HandleFunc("/login", handler.Login)
	r.HandleFunc("/changeAdminPassword", handler.ChangeAdminPassword)
	r.HandleFunc("/getBookingStats", handler.GetBookingStats)
	r.HandleFunc("/getFavoriteItemsStats", handler.GetFavoriteItemsStats)
	r.HandleFunc("/sendTestMail", handler.SendTestMail)
	r.HandleFunc("/getDebts", handler.GetAllCategoryDebts)
	r.HandleFunc("/sendCurrentDebts", handler.SendCurrentDebts)
	r.HandleFunc("/getSettings", handler.GetSettings)
	r.HandleFunc("/updateSetting", handler.UpdateSetting)
	r.HandleFunc("/confirmPaymentIntent", handler.ConfirmPaymentIntent)
	r.HandleFunc("/getStripeCardReader", handler.GetStripeCardReader)
	r.HandleFunc("/cancelReaderAction", handler.CancelReaderAction)

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

	if len(os.Getenv("AVHBS_SSL_CRT_FILE_PATH")) > 0 {
		log.Fatal(server.ListenAndServeTLS(os.Getenv("AVHBS_SSL_CRT_FILE_PATH"), os.Getenv("AVHBS_SSL_KEY_FILE_PATH")))
	} else {
		log.Fatal(server.ListenAndServe())
	}
	log.Fatal(server.ListenAndServe())
}
