package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const PathPrefix = "/api/v1/hello/"

type Greeting struct {
	ID         string `json:"id"`
	Salutation string `json:"salutation"`
}

var Greetings = map[string]string{
	"a": "Hi",
	"b": "Dear Sir or Madam",
	"c": "Moin",
}

func getGreeting(customerID string) string {
	greeting, ok := Greetings[customerID]
	if !ok {
		greeting = "Hello"
	}
	return greeting
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	customerID := strings.TrimPrefix(r.URL.Path, PathPrefix)
	if customerID == "" {
		http.Error(w, "Error: The 'id' parameter is missing from the request URL", http.StatusBadRequest)
		return
	}

	salutation := getGreeting(strings.ToLower(customerID))

	greeting := Greeting{customerID, salutation}
	response, err := json.Marshal(greeting)
	if err != nil {
		http.Error(w, "Error: Unable to serialize greeting to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(response); err != nil {
		log.Printf("Error: Unable to write response: %v", err)
	}
}

func main() {
	http.HandleFunc(PathPrefix, helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
