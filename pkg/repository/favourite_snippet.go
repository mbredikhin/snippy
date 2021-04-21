package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// FavouriteSnippetPostgres structure
type FavouriteSnippetPostgres struct {
	db *sqlx.DB
}

// NewFavouriteSnippetPostgres - FavouriteSnippetPostgres constructor
func NewFavouriteSnippetPostgres(db *sqlx.DB) *FavouriteSnippetPostgres {
	return &FavouriteSnippetPostgres{db: db}
}

// Create - create new favourite snippet
func (r *FavouriteSnippetPostgres) Create(userID, snippetID int) error {
	query := fmt.Sprintf("INSERT INTO %s (snippet_id, user_id) VALUES ($1, $2)", favouriteSnippetsTable)
	_, err := r.db.Exec(query, snippetID, userID)
	return err
}

// GetAll - get all favourite snippet ids
func (r *FavouriteSnippetPostgres) GetAll(userID int) ([]int, error) {
	var ids []int
	query := fmt.Sprintf("SELECT snippet_id FROM %s WHERE user_id=$1", favouriteSnippetsTable)
	err := r.db.Select(&ids, query, userID)
	return ids, err
}

// Delete - remove snippet from favourites
func (r *FavouriteSnippetPostgres) Delete(userID, snippetID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE snippet_id=$1 AND user_id=$2", favouriteSnippetsTable)
	_, err := r.db.Exec(query, snippetID, userID)
	return err
}
