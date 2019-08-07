package outputs

// CheckSessionOutput used to verify session over RPC
type CheckSessionOutput struct {
	Success bool
}

// NewCheckSessionOutput constructor
func NewCheckSessionOutput() *CheckSessionOutput {
	return &CheckSessionOutput{}
}
