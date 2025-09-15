package transactions

import (
	"encoding/json"
	"net/http"

	"github.com/preetamkv/pismo/internal/app/pismo"

	"github.com/go-chi/chi/v5"
)

// Routes defines all routes for Transactions
func Routes(a *pismo.App) http.Handler {
	r := chi.NewRouter()

	r.Post("/", createTransactionHandler(a)) // POST /transactions

	return r

}

// createTransactionHandler handles POST /transactions
func createTransactionHandler(app *pismo.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Validate the incoming request body
		var req CreateTransactionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid JSON", http.StatusBadRequest)
			return
		}
		if err := req.Validate(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add transaction in the DB
		txID, err := createTransaction(app.DB, &req)
		if err != nil {
			http.Error(w, "unable to create Transaction", http.StatusInternalServerError)
		}

		resp := CreateTransactionResponse{
			TransactionID: txID,
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}
