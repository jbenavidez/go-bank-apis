package repository

import (
	"banks/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	InsertUser(user models.User) (int, error)
	AllCustomers() ([]*models.User, error)
}
