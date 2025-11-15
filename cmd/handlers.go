package main

import (
	"banks/models"
	"fmt"
	"net/http"
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
