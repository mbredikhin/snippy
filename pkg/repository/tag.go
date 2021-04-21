package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mbredikhin/snippets"
)

// TagPostgres structure
type TagPostgres struct {
	db *sqlx.DB
}

// NewTagPostgres - TagPostgres constructor
func NewTagPostgres(db *sqlx.DB) *TagPostgres {
	return &TagPostgres{db: db}
}

// Create - create new tag
func (r *TagPostgres) Create(userID int, tag snippets.Tag) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, user_id) VALUES ($1, $2) RETURNING id", tagsTable)
	row := r.db.QueryRow(query, tag.Name, userID)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetAll - get all tags
func (r *TagPostgres) GetAll(userID int) ([]snippets.Tag, error) {
	var tags []snippets.Tag
	query := fmt.Sprintf("SELECT id, name, user_id FROM %s WHERE user_id=$1", tagsTable)
	err := r.db.Select(&tags, query, userID)
	return tags, err
}

// GetByID - get tag by ID
func (r *TagPostgres) GetByID(userID, tagID int) (snippets.Tag, error) {
	var tag snippets.Tag
	query := fmt.Sprintf("SELECT id, name, user_id FROM %s WHERE id=$1 AND user_id=$2", tagsTable)
	err := r.db.Get(&tag, query, tagID, userID)
	return tag, err
}

// Delete - delete tag
func (r *TagPostgres) Delete(userID, tagID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", tagsTable)
	_, err := r.db.Exec(query, tagID, userID)
	return err
}

// Update - update tag
func (r *TagPostgres) Update(userID, tagID int, input snippets.UpdateTagInput) error {
	query := fmt.Sprintf("UPDATE %s t SET name=$1 WHERE t.id=$2 AND t.user_id=$3", tagsTable)
	_, err := r.db.Exec(query, input.Name, tagID, userID)
	return err
}
