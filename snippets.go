package snippets

import "errors"

// List of snippets model
type List struct {
	ID     int    `json:"id" db:"id"`
	UserID string `json:"user_id" db:"user_id"`
	Name   string `json:"name" db:"name" binding:"required"`
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

// Snippet model
type Snippet struct {
	ID       int    `json:"id" db:"id"`
	ListID   int    `json:"list_id" db:"list_id"`
	Name     string `json:"name" db:"name" binding:"required"`
	SyntaxID int    `json:"syntax_id" db:"syntax_id" binding:"required"`
	Content  string `json:"content" db:"content"`
}

// UpdateSnippetInput model
type UpdateSnippetInput struct {
	Name     *string `json:"name"`
	ListID   *int    `json:"list_id"`
	SyntaxID *int    `json:"syntax_id"`
	Content  *string `json:"content"`
}

// Validate - snippet input validation
func (i UpdateSnippetInput) Validate() error {
	if i.Name == nil && i.ListID == nil && i.SyntaxID == nil && i.Content == nil {
		return errors.New("update struture has no values")
	}
	return nil
}

// Tag model
type Tag struct {
	ID     int    `json:"id" db:"id"`
	UserID int    `json:"user_id" db:"user_id"`
	Name   string `json:"name" db:"name" binding:"required"`
}

// UpdateTagInput model
type UpdateTagInput struct {
	Name *string `json:"name"`
}

// Validate - tag input validation
func (i UpdateTagInput) Validate() error {
	if i.Name == nil {
		return errors.New("update struture has no values")
	}
	return nil
}

// Syntax model
type Syntax struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name" binding:"required"`
}
