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

func (app *application) Welcome(w http.ResponseWriter, r *http.Request) {
	resp := JSONResponse{
		Error:   false,
		Message: "Welcome",
	}
	_ = app.writeJSON(w, http.StatusOK, resp)
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

// GetAccountDetails get accounts details
func (app *application) GetAccountDetails(w http.ResponseWriter, r *http.Request) {

	accID, err := strconv.Atoi(chi.URLParam(r, "accId"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	acc, err := app.DB.GetAccount(accID)

	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, acc)
}

// Transactions
func (app *application) AccountTransactions(w http.ResponseWriter, r *http.Request) {

	accID, err := strconv.Atoi(chi.URLParam(r, "accId"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	//
	var transactionRequest TransactionRequest
	err = app.readJSON(w, r, &transactionRequest)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	// validate transaction type
	if strings.ToLower(transactionRequest.TransactionType) != "deposit" && strings.ToLower(transactionRequest.TransactionType) != "withdraw" {
		err := errors.New("invalid transaction type")
		app.errorJSON(w, err)
		return
	}
	//Validate amount
	if transactionRequest.Amount <= 0 {
		err := errors.New("invalid transaction amount")
		app.errorJSON(w, err)
	}
	// get Account
	acc, err := app.DB.GetAccount(accID)

	if err != nil {
		app.errorJSON(w, err)
		return
	}
	switch transactionRequest.TransactionType {
	case "withdraw":
		if acc.Amount < transactionRequest.Amount {
			err := errors.New("insufficient funds")
			app.errorJSON(w, err)
			return
		}
		// apply transaction
		acc.Amount -= transactionRequest.Amount
		acc.UpdatedAt = time.Now()

		err := app.DB.ApplyTransaction(*acc)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

	default:
		// deposit
		acc.Amount += transactionRequest.Amount
		acc.UpdatedAt = time.Now()

		err := app.DB.ApplyTransaction(*acc)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

	}
	resp := JSONResponse{
		Error:   false,
		Message: "Transaction completed",
		Data:    acc,
	}
	_ = app.writeJSON(w, http.StatusOK, resp)
}
