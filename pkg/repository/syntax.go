package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mbredikhin/snippets"
)

// SyntaxPostgres structure
type SyntaxPostgres struct {
	db *sqlx.DB
}

// NewSyntaxPostgres - SyntaxPostgres constructor
func NewSyntaxPostgres(db *sqlx.DB) *SyntaxPostgres {
	return &SyntaxPostgres{db: db}
}

// GetAll - get list of syntaxes
func (r *SyntaxPostgres) GetAll() ([]snippets.Syntax, error) {
	var syntaxList []snippets.Syntax
	query := fmt.Sprintf("SELECT id, name FROM %s", syntaxTable)
	err := r.db.Select(&syntaxList, query)
	return syntaxList, err
}

// Create - create syntax
func (r *SyntaxPostgres) Create(syntax snippets.Syntax) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", syntaxTable)
	row := r.db.QueryRow(query, syntax.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
