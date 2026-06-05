package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Notes struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func landpage_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go")

}

func getNotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path != "/notes" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not suppourted by the server", http.StatusNotFound)
		return
	}

	data, err := os.ReadFile("notes.json")
	if err != nil {
		fmt.Print("there's an error in creating the notes.json file\n", err)

	}
	meow := []Notes{}

	if len(data) > 0 {
		err = json.Unmarshal(data, &meow)
		if err != nil {
			log.Print("There's an error in Unmarshalling the file: \n", err)

		}
	}
	json.NewEncoder(w).Encode(meow)

}
func main() {
	http.HandleFunc("/", landpage_handler)
	http.HandleFunc("/notes", getNotesHandler)
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("The server didn't start due to error")
	}

}
