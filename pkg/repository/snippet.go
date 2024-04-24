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
	query := fmt.Sprintf("INSERT INTO %s (list_id, name, language_id, description, content) VALUES ($1, $2, $3, $4, $5) RETURNING id", snippetsTable)
	row := r.db.QueryRow(query, listID, snippet.Name, snippet.LanguageID, snippet.Description, snippet.Content)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

// GetAll - get all snippets
func (r *SnippetPostgres) GetAll(userID, listID int, tagIDs string, paginationParams *snippets.PaginationParams) ([]snippets.Snippet, error) {
	var snippets []snippets.Snippet
	var err error
	var query string
	if tagIDs == "" {
		query = fmt.Sprintf(`SELECT st.id, st.list_id, st.name, st.language_id, st.description, st.content 
		FROM %s st 
		JOIN %s lt ON st.list_id = lt.id 
		JOIN %s ut ON lt.user_id = ut.id 
		WHERE lt.id=$1 AND ut.id=$2
		LIMIT $3 OFFSET $4`, snippetsTable, listsTable, usersTable)
		err = r.db.Select(&snippets, query, listID, userID, paginationParams.Limit, (paginationParams.Page-1)*paginationParams.Limit)
	} else {
		query = fmt.Sprintf(`SELECT st.id, st.list_id, st.name, st.language_id, st.description, st.content 
		FROM %s st 
		JOIN %s lt ON st.list_id = lt.id
		JOIN %s ut ON lt.user_id = ut.id 
		JOIN %s stt ON st.id = stt.snippet_id 
		WHERE lt.id=$1 AND ut.id=$2 AND stt.tag_id IN ($5)
		LIMIT $3 OFFSET $4`, snippetsTable, listsTable, usersTable, snippetsTagsTable)
		err = r.db.Select(&snippets, query, listID, userID, paginationParams.Limit, (paginationParams.Page-1)*paginationParams.Limit, tagIDs)
	}
	return snippets, err
}

// GetByID - get snippet by ID
func (r *SnippetPostgres) GetByID(userID, snippetID int) (snippets.Snippet, error) {
	var snippet snippets.Snippet
	query := fmt.Sprintf(`SELECT st.id, st.list_id, st.name, st.language_id, st.description, st.content 
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
	SET name=$1, list_id=$2, language_id=$3, description=$4, content=$5
	FROM %s lt, %s ut 
	WHERE st.id=$6 AND st.list_id=lt.id AND lt.user_id=$7 AND lt.user_id=ut.id`, snippetsTable, listsTable, usersTable)
	_, err := r.db.Exec(query, input.Name, input.ListID, input.LanguageID, input.Description, input.Content, snippetID, userID)
	return err
}
