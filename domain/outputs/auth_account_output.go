package outputs

// AuthAccountOutput struct
type AuthAccountOutput struct {
	JWT     string
	Success string
}

// NewAuthAccountOutput constructor
func NewAuthAccountOutput() *AuthAccountOutput {
	return &AuthAccountOutput{}
}
