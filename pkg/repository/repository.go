package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mbredikhin/snippets"
)

// Authorization repo entity interface
type Authorization interface {
	CreateUser(user snippets.User) (int, error)
	GetUser(username, password string) (snippets.User, error)
}

// List repo entity interface
type List interface {
	Create(userID int, list snippets.List) (snippets.List, error)
	GetAll(userID int) ([]snippets.List, error)
	GetByID(userID int, listID int) (snippets.List, error)
	Delete(userID int, listID int) error
	Update(userID int, listID int, input snippets.UpdateListInput) (snippets.List, error)
}

// Snippet repo entity interface
type Snippet interface {
	Create(listID int, snippet snippets.Snippet) (int, error)
	GetAll(userID, listID int) ([]snippets.Snippet, error)
	GetByID(userID, snippetID int) (snippets.Snippet, error)
	Delete(userID, snippetID int) error
	Update(userID, snippetID int, input snippets.UpdateSnippetInput) error
}

// Tag repo entity interface
type Tag interface {
	Create(userID int, tag snippets.Tag) (int, error)
	GetAll(userID int) ([]snippets.Tag, error)
	GetByID(userID, tagID int) (snippets.Tag, error)
	Delete(userID, tagID int) error
	Update(userID, tagID int, input snippets.UpdateTagInput) error
}

// FavouriteSnippet repo entity interface
type FavouriteSnippet interface {
	Create(userID int, snippetID int) error
	GetAll(userID int) ([]int, error)
	Delete(userID, snippetID int) error
}

// Language repo entity interface
type Language interface {
	GetAll() ([]snippets.Language, error)
	Create(snippets.Language) (int, error)
}

// SnippetTag relation interface
type SnippetTag interface {
	Create(userID, snippetID, tagID int) error
	Delete(userID, snippetID, tagID int) error
	GetTagIDs(userID, snippetID int) ([]int, error)
}

// Repository  interface
type Repository struct {
	Authorization
	Snippet
	List
	Tag
	FavouriteSnippet
	Language
	SnippetTag
}

// NewRepository - repository constructor
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:    NewAuthPostgres(db),
		List:             NewListPostgres(db),
		Snippet:          NewSnippetPostgres(db),
		Tag:              NewTagPostgres(db),
		FavouriteSnippet: NewFavouriteSnippetPostgres(db),
		Language:         NewLanguagePostgres(db),
		SnippetTag:       NewSnippetTagPostgres(db),
	}
}
