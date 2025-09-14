package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil) // listen on port 8080
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
