package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	dbP "../database"
)

// GetFeedback forwards API call to get all feedback from database
func GetFeedback(w http.ResponseWriter, r *http.Request) {
	feedbackMap := dbP.GetFeedback()
	response := marshalToJSON(feedbackMap, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// AddFeedback forwards API call to add feedback to database
func AddFeedback(w http.ResponseWriter, r *http.Request) {
	feedback, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	dbP.AddFeedback(string(feedback))
	w.WriteHeader(http.StatusOK)
}

// DeleteFeedback forwards API call to delete feedback from database
func DeleteFeedback(w http.ResponseWriter, r *http.Request) {
	id := UnmarshalInt(r.Body)
	dbP.DeleteFeedback(id)
	w.WriteHeader(http.StatusOK)
}
