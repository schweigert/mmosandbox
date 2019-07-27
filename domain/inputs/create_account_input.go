package inputs

// CreateAccountInput receive some inputs to create an account
type CreateAccountInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
