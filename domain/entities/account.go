package entities

// Account model
type Account struct {
	Model

	Username       string `json:"username"`
	SecurePassword string `json:"secure_password"`
	Email          string `json:"email"`

	Characters []Character `json:"characters"`
}

// NewAccount constructor
func NewAccount() *Account {
	return &Account{}
}
