package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// SnippetTagPostgres structure
type SnippetTagPostgres struct {
	db *sqlx.DB
}

// NewSnippetTagPostgres - SnippetTagPostgres constructor
func NewSnippetTagPostgres(db *sqlx.DB) *SnippetTagPostgres {
	return &SnippetTagPostgres{db: db}
}

// Create - create snippets to tags relation entry
func (r *SnippetTagPostgres) Create(userID, snippetID, tagID int) error {
	query := fmt.Sprintf(`INSERT INTO %s
	SELECT st.id AS snippet_id, tt.id AS tag_id FROM %s st
	JOIN %s lt ON st.list_id = lt.id
	JOIN %s ut ON lt.user_id = ut.id
	JOIN %s tt ON tt.user_id = ut.id
	WHERE ut.id = $1 
		AND st.id = $2 
		AND tt.id = $3 
		AND NOT EXISTS 
			(SELECT snippet_id, tag_id FROM %s stt WHERE stt.snippet_id = $2 AND stt.tag_id = $3)
	`, snippetsTagsTable, snippetsTable, listsTable, usersTable, tagsTable, snippetsTagsTable)
	_, err := r.db.Exec(query, userID, snippetID, tagID)
	return err
}

// Delete - delete snippets to tags relation entry
func (r *SnippetTagPostgres) Delete(userID, snippetID, tagID int) error {
	query := fmt.Sprintf(`DELETE FROM %s stt 
		WHERE snippet_id IN (
			SELECT st.id FROM %s st
			JOIN %s lt ON st.list_id = lt.id
			JOIN %s ut ON lt.user_id = ut.id
			WHERE ut.id = $1 AND st.id = $2)
		AND tag_id IN (
			SELECT tt.id FROM %s tt
			JOIN %s ut ON tt.user_id = ut.id
			WHERE ut.id = $1 AND tt.id = $3)`, snippetsTagsTable, snippetsTable, listsTable, usersTable, tagsTable, usersTable)
	_, err := r.db.Exec(query, userID, snippetID, tagID)
	return err
}

// GetTagIDs - get tag ids by snippet id
func (r *SnippetTagPostgres) GetTagIDs(userID, snippetID int) ([]int, error) {
	var ids []int
	query := fmt.Sprintf(`SELECT stt.tag_id FROM %s stt
	JOIN %s st ON st.id = stt.snippet_id
	JOIN %s lt ON st.list_id = lt.id
	JOIN %s ut ON lt.user_id = ut.id
	WHERE ut.id = $1 AND st.id = $2`, snippetsTagsTable, snippetsTable, listsTable, usersTable)
	err := r.db.Select(&ids, query, userID, snippetID)
	return ids, err
}
