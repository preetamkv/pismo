package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/preetamkv/pismo/internal/app/pismo"

	"github.com/go-chi/chi/v5"
)

// Routes defines all routes for Accounts
func Routes(a *pismo.App) http.Handler {
	r := chi.NewRouter()

	r.Post("/", createAccountHandler(a)) // GET /accounts
	r.Get("/{id}", getAccountHandler(a)) // GET /accounts/{id}

	return r

}

// createAccountHandler handles POST /accounts
func createAccountHandler(app *pismo.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate the incoming request body
		var req CreateAccountRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		if err := req.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add Account in the DB
		accID, err := createAccount(app.DB, &req)
		if err != nil {
			http.Error(w, "Unable to create account", http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Created Account %s\n", accID)
	}
}

// getAccountHandler handles GET /accounts/{id}
func getAccountHandler(app *pismo.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		acc, err := fetchAccount(app.DB, id)
		if err != nil {
			http.Error(w, "invalid JSON", http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "List account %v\n", *acc)
	}
}
