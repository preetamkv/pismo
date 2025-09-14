package accounts

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", createAccount) // GET /accounts
	r.Get("/{id}", getAccount) // GET /accounts/{id}

	return r

}

func createAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Creating Account")
}

func getAccount(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "List account %s\n", id)
}
