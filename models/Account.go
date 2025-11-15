package models

import "time"

type Account struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	AccountType string    `json:"account_type"`
	Amount      float64   `json:"amount"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
