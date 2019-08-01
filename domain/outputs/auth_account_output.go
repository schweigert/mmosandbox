package outputs

// AuthAccountOutput struct
type AuthAccountOutput struct {
	JWT     string
	Success bool
}

// NewAuthAccountOutput constructor
func NewAuthAccountOutput() *AuthAccountOutput {
	return &AuthAccountOutput{}
}
