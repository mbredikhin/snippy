package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mbredikhin/snippets"
)

// SnippetPostgres structure
type SnippetPostgres struct {
	db *sqlx.DB
}

// NewSnippetPostgres - SnippetPostgres constructor
func NewSnippetPostgres(db *sqlx.DB) *SnippetPostgres {
	return &SnippetPostgres{db: db}
}

// Create - create new snippet
func (r *SnippetPostgres) Create(listID int, snippet snippets.Snippet) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (list_id, name, language_id, content) VALUES ($1, $2, $3, $4) RETURNING id", snippetsTable)
	row := r.db.QueryRow(query, listID, snippet.Name, snippet.LanguageID, snippet.Content)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetAll - get all snippets
func (r *SnippetPostgres) GetAll(userID, listID int) ([]snippets.Snippet, error) {
	var snippets []snippets.Snippet
	query := fmt.Sprintf(`SELECT st.id, st.list_id, st.name, st.language_id, st. content 
	FROM %s st 
	JOIN %s lt ON st.list_id = lt.id 
	JOIN %s ut ON lt.user_id = ut.id 
	WHERE lt.id=$1 AND ut.id=$2`, snippetsTable, listsTable, usersTable)
	err := r.db.Select(&snippets, query, listID, userID)
	return snippets, err
}

// GetByID - get snippet by ID
func (r *SnippetPostgres) GetByID(userID, snippetID int) (snippets.Snippet, error) {
	var snippet snippets.Snippet
	query := fmt.Sprintf(`SELECT st.id, st.list_id, st.name, st.language_id, st.content 
	FROM %s st 
	JOIN %s lt ON st.list_id = lt.id 
	JOIN %s ut ON lt.user_id = ut.id 
	WHERE st.id=$1 AND ut.id=$2`, snippetsTable, listsTable, usersTable)
	err := r.db.Get(&snippet, query, snippetID, userID)
	return snippet, err
}

// Delete - delete snippet
func (r *SnippetPostgres) Delete(userID, snippetID int) error {
	query := fmt.Sprintf(`DELETE FROM %s st USING %s lt 
	WHERE st.id=$1 AND st.list_id=lt.id AND lt.user_id=$2`, snippetsTable, listsTable)
	_, err := r.db.Exec(query, snippetID, userID)
	return err
}

// Update - update snippet
func (r *SnippetPostgres) Update(userID, snippetID int, input snippets.UpdateSnippetInput) error {
	query := fmt.Sprintf(`UPDATE %s st 
	SET name=$1, list_id=$2, language_id=$3, content=$4 
	FROM %s lt, %s ut 
	WHERE st.id=$5 AND st.list_id=lt.id AND lt.user_id=$6 AND lt.user_id=ut.id`, snippetsTable, listsTable, usersTable)
	_, err := r.db.Exec(query, input.Name, input.ListID, input.LanguageID, input.Content, snippetID, userID)
	return err
}
