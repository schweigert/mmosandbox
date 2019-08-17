package outputs

import "github.com/schweigert/mmosandbox/domain/entities"

// ListenMessagesOutput struct
type ListenMessagesOutput struct {
	StartSessionOutput
	Messages []*entities.Message
}

// NewListenMessagesOutput constructor
func NewListenMessagesOutput() *ListenMessagesOutput {
	return &ListenMessagesOutput{}
}
