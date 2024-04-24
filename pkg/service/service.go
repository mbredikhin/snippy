package service

import (
	"github.com/mbredikhin/snippets"
	"github.com/mbredikhin/snippets/pkg/repository"
)

// Authorization - authorization service interface
type Authorization interface {
	CreateUser(user snippets.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, int64, error)
	BlacklistToken(token string, expiresAt int64) error
	CheckIfTokenBlacklisted(token string) bool
	RemoveExpiredTokensFromBlacklist(timestamp int64) error
}

// List service interface
type List interface {
	Create(userID int, list snippets.List) (snippets.List, error)
	GetAll(userID int, paginationParams *snippets.PaginationParams) ([]snippets.List, error)
	GetByID(userID int, listID int) (snippets.List, error)
	Delete(userID int, listID int) error
	Update(userID int, listID int, input snippets.UpdateListInput) (snippets.List, error)
}

// Snippet service interface
type Snippet interface {
	Create(listID int, snippet snippets.Snippet) (int, error)
	GetAll(userID, listID int, tagIDs string, paginationParams *snippets.PaginationParams) ([]snippets.Snippet, error)
	GetByID(userID, snippetID int) (snippets.Snippet, error)
	Delete(userID, snippetID int) error
	Update(userID, snippetID int, input snippets.UpdateSnippetInput) error
	GetFavourites(userID int) ([]int, error)
	AddFavourite(userID, snippetID int) error
	RemoveFavourite(userID, snippetID int) error
	AddTag(userID, snippetID, tagID int) error
	RemoveTag(userID, snippetID, tagID int) error
	GetTagIDs(userID, snippetID int) ([]int, error)
	GetLanguages() ([]snippets.Language, error)
	CreateLanguage(snippets.Language) (int, error)
}

// Tag service interface
type Tag interface {
	Create(userID int, tag snippets.Tag) (int, error)
	GetAll(userID int) ([]snippets.Tag, error)
	GetByID(userID, tagID int) (snippets.Tag, error)
	Delete(userID, tagID int) error
	Update(userID, tagID int, input snippets.UpdateTagInput) error
}

// Service - Service struct
type Service struct {
	Authorization
	List
	Snippet
	Tag
}

// NewService - Service constructor
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		List:          NewListService(repos.List),
		Snippet:       NewSnippetService(repos.Snippet, repos.SnippetTag, repos.FavouriteSnippet, repos.Language),
		Tag:           NewTagService(repos.Tag, repos.SnippetTag),
	}
}
