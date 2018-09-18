package api

// User data structure
type User struct {
	ID       string
	Name     string
	Email    string
	Birthday string
	Gender   string
	Avatar   string
	Error    error
}
