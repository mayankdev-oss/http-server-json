package main

import (
	"fmt"
	"net/http"
)

func landpage_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go")

}
func main() {
	http.HandleFunc("/", landpage_handler)
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("The server didn't start due to error")
	}

}
