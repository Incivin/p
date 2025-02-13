package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"p/internal/database"
	"p/internal/handlers"
	"p/internal/taskService"
)

func main() {
	database.InitDB()
	database.DB.AutoMigrate(&taskService.Task{})
	repo := taskService.NewTaskRepository(database.DB)
	service := taskService.NewService(repo)
	handler := handlers.NewHandler(service)

	router := mux.NewRouter()

	router.HandleFunc("/api/get", handler.GetTasksHandler).Methods("GET")
	router.HandleFunc("/api/post", handler.PostTaskHandler).Methods("POST")
	router.HandleFunc("/api/patch/{id}", handler.PatchTaskHandler).Methods("PATCH")
	router.HandleFunc("/api/delete/{id}", handler.DeleteTaskHandler).Methods("DELETE")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("Error starting server:", err)
	}

}
