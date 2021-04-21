package service

import (
	"github.com/mbredikhin/snippets"
	"github.com/mbredikhin/snippets/pkg/repository"
)

// TagService - Tag service
type TagService struct {
	tagRepo        repository.Tag
	snippetTagRepo repository.SnippetTag
}

// NewTagService - Tag service constructor
func NewTagService(tagRepo repository.Tag, snippetTagRepo repository.SnippetTag) *TagService {
	return &TagService{tagRepo, snippetTagRepo}
}

// Create - Create new tag
func (s *TagService) Create(userID int, tag snippets.Tag) (int, error) {
	return s.tagRepo.Create(userID, tag)
}

// GetAll - Get all the user's tags
func (s *TagService) GetAll(userID int) ([]snippets.Tag, error) {
	return s.tagRepo.GetAll(userID)
}

// GetByID - Get tag
func (s *TagService) GetByID(userID, tagID int) (snippets.Tag, error) {
	return s.tagRepo.GetByID(userID, tagID)
}

// Delete - Delete tag
func (s *TagService) Delete(userID, tagID int) error {
	return s.tagRepo.Delete(userID, tagID)
}

// Update - Edit tag
func (s *TagService) Update(userID, tagID int, input snippets.UpdateTagInput) error {
	return s.tagRepo.Update(userID, tagID, input)
}
