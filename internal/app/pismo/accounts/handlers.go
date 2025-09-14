package accounts

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/preetamkv/pismo/internal/app/pismo"
)

func Routes(a *pismo.App) http.Handler {
	r := chi.NewRouter()

	r.Post("/", createAccount(a)) // GET /accounts
	r.Get("/{id}", getAccount(a)) // GET /accounts/{id}

	return r

}

func createAccount(app *pismo.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Creating Account")
	}
}

func getAccount(app *pismo.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		fmt.Fprintf(w, "List account %s\n", id)
	}
}
