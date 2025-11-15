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

func (m *PostgresDBRepo) AllCustomers() ([]*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select
			first_name, last_name, email, username
		from
			users
	`
	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customerList []*models.User

	for rows.Next() {
		var customer models.User
		err := rows.Scan(
			&customer.FirstName,
			&customer.LastName,
			&customer.Email,
			&customer.Username,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		customerList = append(customerList, &customer)
	}
	return customerList, nil
}

func (m *PostgresDBRepo) Getuser(userID int) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()
	query := `
		select
			first_name, last_name, email, username
		from
			users
		where id = $1
	`
	row := m.DB.QueryRowContext(ctx, query, userID)
	var user models.User
	err := row.Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Username,
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil

}

func (m *PostgresDBRepo) UpdateUser(userID int, userObj models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `
		update users set first_name=$1, last_name=$2 , email = $3, username = $4  
		where id = $5
	`
	_, err := m.DB.ExecContext(ctx, stmt,
		userObj.FirstName,
		userObj.LastName,
		userObj.Email,
		userObj.Username,
		userID,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
