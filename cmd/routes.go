package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {

	mux := chi.NewRouter()

	mux.Get("/welcome", a.Welcome)
	// Customer APIs
	mux.Post("/create-customer", a.CreateUser)
	mux.Get("/customers", a.AllCustomers)
	mux.Get("/customers/{id}", a.GetCustomer)
	mux.Put("/customers/{id}", a.UpdateCustomer)
	// TODO: account
	mux.Post("/customers/{id}/account", a.CreateAccount)
	return mux
}
