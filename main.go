package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

func CreateMessage(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		fmt.Fprintln(w, "Json error")
		return
	}
	if err := DB.Create(&msg).Error; err != nil {
		fmt.Fprintln(w, "Fail add DB")
		return
	}
	fmt.Fprintln(w, "Записано в message", msg.Task)
}
func GetMessages(w http.ResponseWriter, r *http.Request) {
	var message []Message
	if err := DB.Find(&message).Error; err != nil {
		fmt.Fprintln(w, "Fail find DB")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

func main() {
	InitDB()
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	http.ListenAndServe(":8080", router)

}
