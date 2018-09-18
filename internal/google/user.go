package google

// User google data structure
type User struct {
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Gender        string `json:"gender"`
	Email         string `json:"email"`
	ID            string `json:"id"`
	Error         *Error `json:"error,omitempty"`
}

// Error data structure
type Error struct {
	Errors  interface{} `json:"errors"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
}
