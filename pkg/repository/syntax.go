package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mbredikhin/snippets"
)

// LanguagePostgres structure
type LanguagePostgres struct {
	db *sqlx.DB
}

// NewLanguagePostgres - LanguagePostgres constructor
func NewLanguagePostgres(db *sqlx.DB) *LanguagePostgres {
	return &LanguagePostgres{db: db}
}

// GetAll - get list of languages
func (r *LanguagePostgres) GetAll() ([]snippets.Language, error) {
	var languages []snippets.Language
	query := fmt.Sprintf("SELECT id, name FROM %s", languagesTable)
	err := r.db.Select(&languages, query)
	return languages, err
}

// Create - create language
func (r *LanguagePostgres) Create(language snippets.Language) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", languagesTable)
	row := r.db.QueryRow(query, language.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
