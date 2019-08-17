package inputs

// ChatInput struct
type ChatInput struct {
	CharacterID int    `json:"character_id"`
	Body        string `json:"body"`

	CheckSessionInput *CheckSessionInput `json:"check_session_input"`
}

// HasBody verify
func (in *ChatInput) HasBody() bool {
	return in.Body != ""
}

// NewChatInput constructor
func NewChatInput() *ChatInput {
	return &ChatInput{}
}
