package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable             = "users"
	snippetsTable          = "snippets"
	listsTable             = "lists"
	tagsTable              = "tags"
	favouriteSnippetsTable = "favourite_snippets"
	snippetsTagsTable      = "snippets_tags"
	languagesTable         = "languages"
)

// Config db
type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

// NewPostgresDB - postgres connection
func NewPostgresDB(conf Config) (*sqlx.DB, error) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
