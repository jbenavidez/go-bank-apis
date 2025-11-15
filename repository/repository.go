package repository

import (
	"banks/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	InsertUser(user models.User) (int, error)
	AllCustomers() ([]*models.User, error)
	Getuser(userID int) (*models.User, error)
	UpdateUser(userID int, userObj models.User) error
	InsertAccount(account models.Account) (int, error)
}
