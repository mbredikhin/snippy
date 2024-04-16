package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mbredikhin/snippets"
)

// ListPostgres structure
type ListPostgres struct {
	db *sqlx.DB
}

// NewListPostgres - ListPostgres constructor
func NewListPostgres(db *sqlx.DB) *ListPostgres {
	return &ListPostgres{db: db}
}

// Create - create new list
func (r *ListPostgres) Create(userID int, list snippets.List) (snippets.List, error) {
	query := fmt.Sprintf("INSERT INTO %s (user_id, name) VALUES ($1, $2) RETURNING id", listsTable)
	row := r.db.QueryRow(query, userID, list.Name)
	if err := row.Scan(&list.ID); err != nil {
		return list, err
	}
	return list, nil
}

// GetAll - get all lists
func (r *ListPostgres) GetAll(userID int) ([]snippets.List, error) {
	var lists []snippets.List
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE user_id=$1", listsTable)
	err := r.db.Select(&lists, query, userID)
	return lists, err
}

// GetByID - get list by ID
func (r *ListPostgres) GetByID(userID int, listID int) (snippets.List, error) {
	var list snippets.List
	query := fmt.Sprintf("SELECT id, name FROM %s WHERE user_id=$1 AND id=$2", listsTable)
	err := r.db.Get(&list, query, userID, listID)
	return list, err
}

// Delete - delete list
func (r *ListPostgres) Delete(userID int, listID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 AND id=$2", listsTable)
	_, err := r.db.Exec(query, userID, listID)
	return err
}

// Update - update list
func (r *ListPostgres) Update(userID int, listID int, input snippets.UpdateListInput) (snippets.List, error) {
	query := fmt.Sprintf("UPDATE %s t SET name=$1 WHERE t.id=$2 AND t.user_id=$3", listsTable)
	_, err := r.db.Exec(query, input.Name, listID, userID)
	list := snippets.List{ID: listID, Name: *input.Name}
	return list, err
}
