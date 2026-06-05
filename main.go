package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("The server didn't start due to error")
	}

}
