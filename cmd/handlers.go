package main

import (
	"banks/models"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

type Payload struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func (app *application) Welcome(w http.ResponseWriter, r *http.Request) {
	var payload = Payload{
		Status:  "active",
		Message: "welcome to my Bank APIs :)",
		Version: "2.0",
	}
	_ = app.writeJSON(w, http.StatusOK, payload)
}

// CreateUser create user record
func (app *application) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload models.User
	err := app.readJSON(w, r, &payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	// inser user on db
	userID, err := app.DB.InsertUser(payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user_id ", userID)

}

// AllCustomers  get all customer
func (app *application) AllCustomers(w http.ResponseWriter, r *http.Request) {

	// Get customers
	customerList, err := app.DB.AllCustomers()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("customers list retrieved")

	_ = app.writeJSON(w, http.StatusOK, customerList)

}

// GetCustomer get customer details
func (app *application) GetCustomer(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	c, err := app.DB.Getuser(userID)

	if err != nil {
		app.errorJSON(w, err)
		return
	}
	//return customer
	_ = app.writeJSON(w, http.StatusOK, c)
}

// UpdateCustomer update customer info
func (app *application) UpdateCustomer(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	// Get payload
	var payload models.User
	err = app.readJSON(w, r, &payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	// get customer record
	_, err = app.DB.Getuser(userID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	// update user
	err = app.DB.UpdateUser(userID, payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	fmt.Println(userID)

	resp := JSONResponse{
		Error:   false,
		Message: "Customer updated",
	}
	_ = app.writeJSON(w, http.StatusOK, resp)
}

// CreateAccount create an account for a given user
func (app *application) CreateAccount(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	var account models.Account
	err = app.readJSON(w, r, &account)

	if err != nil {
		fmt.Println(err)
		app.errorJSON(w, err)
		return
	}
	// check if amount is greater than 0
	if account.Amount < 1 {
		err := errors.New("account amount should be greater than zero")
		app.errorJSON(w, err)

	}

	// validate account type
	if strings.ToLower(account.AccountType) != "checking" && strings.ToLower(account.AccountType) != "saving" {
		err := errors.New("invalid account type")
		app.errorJSON(w, err)
	}
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()
	account.UserID = userID
	//create account
	_, err = app.DB.InsertAccount(account)
	if err != nil {
		fmt.Println(err)
		app.errorJSON(w, err)
		return
	}
	//

	//return response
	resp := JSONResponse{
		Error:   false,
		Message: "Account created",
	}
	_ = app.writeJSON(w, http.StatusAccepted, resp)
}

// GetUserAccounts gets all account for a given user
func (app *application) GetUserAccounts(w http.ResponseWriter, r *http.Request) {

	userID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	userAccounts, err := app.DB.GetAccountsByUserId(userID)
	if err != nil {
		fmt.Println(err)
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, userAccounts)
}
