package outputs

// StartSessionOutput struct
type StartSessionOutput struct {
	JWT     string
	Success bool
}

// NewStartSessionOutput constructor
func NewStartSessionOutput() *StartSessionOutput {
	return &StartSessionOutput{}
}
