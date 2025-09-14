package transactions

import (
	"fmt"
	"net/http"

	"github.com/preetamkv/pismo/internal/app/pismo"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/preetamkv/pismo/internal/pkg/models"
)

func Routes(a *pismo.App) http.Handler {
	r := chi.NewRouter()

	r.Post("/", createTransaction(a)) // POST /transactions

	return r

}

func createTransaction(app *pismo.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		txID := uuid.New()
		tx := models.Transaction{
			TransactionID: txID.String(),
		}
		fmt.Fprintf(w, "Transaction created %s\n", tx.TransactionID)
	}
}
