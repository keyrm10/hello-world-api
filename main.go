package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Greeting struct {
	ID         string `json:"id"`
	Salutation string `json:"salutation"`
}

var Greetings = map[string]string{
	"A": "Hi",
	"B": "Dear Sir or Madam",
	"C": "Moin",
}

func getGreeting(customerID string) string {
	greeting, ok := Greetings[customerID]
	if !ok {
		greeting = "Hello"
	}
	return greeting
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	customerID := r.URL.Query().Get("id")
	if customerID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	salutation := getGreeting(customerID)

	greeting := Greeting{customerID, salutation}
	response, err := json.Marshal(greeting)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
