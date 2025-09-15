package accounts

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/preetamkv/pismo/internal/app/pismo"

	"github.com/go-chi/chi/v5"
)

// Routes defines all routes for Accounts
func Routes(a *pismo.App) http.Handler {
	r := chi.NewRouter()

	r.Post("/", createAccountHandler(a)) // POST /accounts
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
			if strings.Contains(err.Error(), "duplicate key value") {
				http.Error(w, "Duplicate Document Number", http.StatusConflict)
				return
			}
			http.Error(w, "Unable to create account", http.StatusInternalServerError)
			return
		}
		resp := CreateAccountResponse{
			AccountNumber: accID,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}

// getAccountHandler handles GET /accounts/{id}
func getAccountHandler(app *pismo.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		acc, err := FetchAccount(app.DB, id)
		if err != nil {
			if strings.Contains(err.Error(), "record not found") {
				http.Error(w, "Account doesnt exist", http.StatusNotFound)
				return
			}
			http.Error(w, "failed to fetch account: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(acc); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}
