package inputs

// CheckSessionInput struct
type CheckSessionInput struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

// NewCheckSessionInput struct
func NewCheckSessionInput() *CheckSessionInput {
	return &CheckSessionInput{}
}
