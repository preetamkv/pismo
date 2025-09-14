package main

import (
	"fmt"
	"net/http"

	"github.com/preetamkv/pismo/internal/app/pismo/accounts"
	"github.com/preetamkv/pismo/internal/app/pismo/transactions"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Mount("/accounts", accounts.Routes())
	r.Mount("/transactions", transactions.Routes())

	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil) // listen on port 8080
	if err != nil {
		fmt.Println("Error starting server:", r)
	}
}
