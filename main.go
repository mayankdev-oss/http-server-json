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

func NotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.URL.Path != "/notes" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
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
	case "POST":
		var newnote Notes
		err := json.NewDecoder(r.Body).Decode(&newnote)
		if err != nil {
			fmt.Println("Couldn't decode the incoming request")
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
		newnote.ID = len(meow) + 1
		meow = append(meow, newnote)
		date, errr := json.MarshalIndent(meow, "", " ")
		if errr != nil {
			fmt.Println("Error in marshalling the data into a json")
		}
		eror := os.WriteFile("notes.json", date, 0644)
		if eror != nil {
			fmt.Println("Error in writing the file in disk")
		}
		w.WriteHeader(http.StatusCreated)

	default:
		http.Error(w, "This Method isn't suppourted for now", http.StatusMethodNotAllowed)
		return
	}

}
func main() {
	http.HandleFunc("/", landpage_handler)
	http.HandleFunc("/notes", NotesHandler)
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("The server didn't start due to error")
	}

}
