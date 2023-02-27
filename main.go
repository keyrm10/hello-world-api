package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

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
	urlParts := strings.Split(r.URL.Path, "/")
	customerID := urlParts[len(urlParts)-1]	
	if customerID == "" {
		http.Error(w, "Missing id parameter", http.StatusBadRequest)
		return
	}

	salutation := getGreeting(strings.ToLower(customerID))

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
	http.HandleFunc("/hello/", helloHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
