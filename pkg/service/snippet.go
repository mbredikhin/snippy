package service

import (
	"github.com/mbredikhin/snippets"
	"github.com/mbredikhin/snippets/pkg/repository"
)

// SnippetService - Snippet service
type SnippetService struct {
	snippetRepo          repository.Snippet
	snippetTagRepo       repository.SnippetTag
	favouriteSnippetRepo repository.FavouriteSnippet
	syntaxRepo           repository.Syntax
}

// NewSnippetService - Snippet service constructor
func NewSnippetService(
	snippetRepo repository.Snippet,
	snippetTagRepo repository.SnippetTag,
	favouriteSnippetRepo repository.FavouriteSnippet,
	syntaxRepo repository.Syntax) *SnippetService {
	return &SnippetService{snippetRepo, snippetTagRepo, favouriteSnippetRepo, syntaxRepo}
}

// Create - Create new snippet
func (s *SnippetService) Create(listID int, snippet snippets.Snippet) (int, error) {
	return s.snippetRepo.Create(listID, snippet)
}

// GetAll - Get all snippets from list
func (s *SnippetService) GetAll(userID, listID int) ([]snippets.Snippet, error) {
	return s.snippetRepo.GetAll(userID, listID)
}

// GetByID - Get snippet
func (s *SnippetService) GetByID(userID, snippetID int) (snippets.Snippet, error) {
	return s.snippetRepo.GetByID(userID, snippetID)
}

// Delete - Delete snippet
func (s *SnippetService) Delete(userID, snippetID int) error {
	return s.snippetRepo.Delete(userID, snippetID)
}

// Update - Edit snippet
func (s *SnippetService) Update(userID, snippetID int, input snippets.UpdateSnippetInput) error {
	return s.snippetRepo.Update(userID, snippetID, input)
}

// GetFavourites - Get all the user's favourite snippets
func (s *SnippetService) GetFavourites(userID int) ([]int, error) {
	return s.favouriteSnippetRepo.GetAll(userID)
}

// AddFavourite - Add snippet to favourites
func (s *SnippetService) AddFavourite(userID, snippetID int) error {
	return s.favouriteSnippetRepo.Create(userID, snippetID)
}

// RemoveFavourite - Remove snippet from favourites
func (s *SnippetService) RemoveFavourite(userID, snippetID int) error {
	return s.favouriteSnippetRepo.Delete(userID, snippetID)
}

// AddTag - Add tag to snippet
func (s *SnippetService) AddTag(userID, snippetID, tagID int) error {
	return s.snippetTagRepo.Create(userID, snippetID, tagID)
}

// RemoveTag - Remove tag from snippet
func (s *SnippetService) RemoveTag(userID, snippetID, tagID int) error {
	return s.snippetTagRepo.Delete(userID, snippetID, tagID)
}

// GetTagIDs - Get ID's of all the snippet tags
func (s *SnippetService) GetTagIDs(userID, snippetID int) ([]int, error) {
	return s.snippetTagRepo.GetTagIDs(userID, snippetID)
}

// GetSyntaxList - Get syntax list
func (s *SnippetService) GetSyntaxList() ([]snippets.Syntax, error) {
	return s.syntaxRepo.GetAll()
}

// CreateSyntax - Create new syntax
func (s *SnippetService) CreateSyntax(syntax snippets.Syntax) (int, error) {
	return s.syntaxRepo.Create(syntax)
}
