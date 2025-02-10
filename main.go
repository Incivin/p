package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)

}
func GetMessages(w http.ResponseWriter, r *http.Request) {
	var msg []Message
	if err := DB.Find(&msg).Error; err != nil {
		fmt.Fprintln(w, "Fail find DB")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var msg Message
	if err := DB.First(&msg, id).Error; err != nil {
		fmt.Fprintln(w, "Message not found")
		return
	}
	var updateData Message
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		fmt.Fprintln(w, "Invalid JSON")
		return
	}
	if updateData.Task != "" {
		msg.Task = updateData.Task
	}
	msg.IsDone = updateData.IsDone

	if err := DB.Save(&msg).Error; err != nil {
		fmt.Fprintln(w, "Fail to update task")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)

}
func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var msg Message
	if err := DB.First(&msg, id).Error; err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}
	if err := DB.Delete(&msg).Error; err != nil {
		http.Error(w, "Fail to delete task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)

}

func main() {
	InitDB()
	DB.AutoMigrate(&Message{})
	router := mux.NewRouter()
	router.HandleFunc("/api/messages/{id}", UpdateMessage).Methods("PATCH")
	router.HandleFunc("/api/messages/{id}", DeleteMessage).Methods("DELETE")
	router.HandleFunc("/api/messages", CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages", GetMessages).Methods("GET")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Error starting server:", err)
	}

}
