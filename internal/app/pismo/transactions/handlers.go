package transactions

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", createTransaction) // POST /transactions

	return r

}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating tx")
}
