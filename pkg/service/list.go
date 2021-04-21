package service

import (
	"github.com/mbredikhin/snippets"
	"github.com/mbredikhin/snippets/pkg/repository"
)

// ListService - list service
type ListService struct {
	repo repository.List
}

// NewListService - List service constructor
func NewListService(repo repository.List) *ListService {
	return &ListService{repo: repo}
}

// Create - Create new list
func (l *ListService) Create(userID int, list snippets.List) (int, error) {
	return l.repo.Create(userID, list)
}

// GetAll - Get all user's lists
func (l *ListService) GetAll(userID int) ([]snippets.List, error) {
	return l.repo.GetAll(userID)
}

// GetByID - Get list
func (l *ListService) GetByID(userID int, listID int) (snippets.List, error) {
	return l.repo.GetByID(userID, listID)
}

// Delete - Delete list
func (l *ListService) Delete(userID int, listID int) error {
	return l.repo.Delete(userID, listID)
}

// Update - Edit list
func (l *ListService) Update(userID int, listID int, input snippets.UpdateListInput) error {
	return l.repo.Update(userID, listID, input)
}
