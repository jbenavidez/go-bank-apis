package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() http.Handler {

	mux := chi.NewRouter()

	mux.Get("/welcome", a.Welcome)
	// Customers APIs
	mux.Post("/create-customer", a.CreateUser)
	mux.Get("/customers", a.AllCustomers)
	mux.Get("/customers/{id}", a.GetCustomer)
	mux.Put("/customers/{id}", a.UpdateCustomer)
	// Accounts APIs
	mux.Post("/customers/{id}/accounts", a.CreateAccount)
	mux.Get("/customers/{id}/accounts", a.GetUserAccounts)
	mux.Get("/customers/{id}/accounts/{accId}", a.GetAccountDetails)
	return mux
}
