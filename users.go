package snippets

// User model
type User struct {
	ID       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
