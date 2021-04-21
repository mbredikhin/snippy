package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mbredikhin/snippets"
)

// AuthPostgres - db authentication
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres - db authentication constructor
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser - create new user in db
func (r *AuthPostgres) CreateUser(user snippets.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetUser - get user from db
func (r *AuthPostgres) GetUser(username, password string) (snippets.User, error) {
	var user snippets.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
