package dbrepo

import (
	"banks/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) InsertUser(user models.User) (int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	stmt := `
		insert into Users
			(first_name, last_name, email, username)
		values	
			($1,$2,$3,$4) 
		returning id
	`

	var newID int

	err := m.DB.QueryRowContext(ctx, stmt,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Username,
	).Scan(&newID)

	if err != nil {
		fmt.Println("err creating user", err)
		return 0, err
	}
	fmt.Println("user was created", newID)
	return newID, nil

}
