package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	rBody := requestBody{}
	json.NewDecoder(r.Body).Decode(&rBody)
	task = rBody.Message
}
func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello", task)
}

func main() {
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/post", PostHandler)
	http.ListenAndServe("localhost:8080", nil)

}
