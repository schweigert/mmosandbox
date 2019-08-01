package outputs

// StartSessionOutput struct
type StartSessionOutput struct {
	Token   string
	Success bool
}

// NewStartSessionOutput constructor
func NewStartSessionOutput() *StartSessionOutput {
	return &StartSessionOutput{}
}
