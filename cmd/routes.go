package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {

	mux := chi.NewRouter()

	mux.Get("/welcome", a.Welcome)
	mux.Post("/create-customer", a.CreateUser)
	mux.Get("/customers", a.AllCustomers)
	return mux
}
