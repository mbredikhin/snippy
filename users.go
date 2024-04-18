package snippets

// User model
type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignInInput model
type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SignInResponse model
type SignInResponse struct {
	Token *string `json:"token"`
}
