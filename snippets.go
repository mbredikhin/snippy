package snippets

import (
	"errors"
)

// List of snippets model
type List struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

// UpdateListInput model
type UpdateListInput struct {
	Name *string `json:"name"`
}

// Validate - list input validation
func (i UpdateListInput) Validate() error {
	if i.Name == nil {
		return errors.New("update struture has no values")
	}
	return nil
}

// DeleteListResponse model
type DeleteListResponse struct {
	ID *int `json:"id"`
}

// Snippet model
type Snippet struct {
	ID         int    `json:"id" db:"id"`
	ListID     int    `json:"list_id" db:"list_id"`
	Name       string `json:"name" db:"name" binding:"required"`
	LanguageID int    `json:"language_id" db:"language_id" binding:"required"`
	Content    string `json:"content" db:"content" binding:"required"`
}

// CreateSnippetRespnse model
type CreateSnippetResponse struct {
	ID *int `json:"id"`
}

// UpdateSnippetInput model
type UpdateSnippetInput struct {
	Name       *string `json:"name"`
	ListID     *int    `json:"list_id"`
	LanguageID *int    `json:"language_id"`
	Content    *string `json:"content"`
}

// AddSnippetToFavouritesInput model
type AddSnippetToFavouritesInput struct {
	ID *int `json:"id"`
}

// RemoveSnippetFromFavouritesInput model
type RemoveSnippetFromFavouritesInput struct {
	ID *int `json:"id"`
}

// Validate - snippet input validation
func (i UpdateSnippetInput) Validate() error {
	if i.Name == nil && i.ListID == nil && i.LanguageID == nil && i.Content == nil {
		return errors.New("update struture has no values")
	}
	return nil
}

// Tag model
type Tag struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

// UpdateTagInput model
type UpdateTagInput struct {
	Name *string `json:"name"`
}

// AddTagResponse model
type AddTagResponse struct {
	ID *int `json:"id"`
}

// Validate - tag input validation
func (i UpdateTagInput) Validate() error {
	if i.Name == nil {
		return errors.New("update struture has no values")
	}
	return nil
}

// Language model
type Language struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}

// AddLanguageResponse model
type AddLanguageResponse struct {
	ID *int `json:"id"`
}

// AddTagToSnippetInput model
type AddTagToSnippetInput struct {
	TagID *int `json:"tag_id"`
}

// RemoveTagFromSnippetInput model
type RemoveTagFromSnippetInput struct {
	TagID *int `json:"tag_id"`
}
