package inputs

// ChatInput struct
type ChatInput struct {
	CharacterID int `json:"character_id"`
	Body        int `json:"body"`

	CheckSessionInput *CheckSessionInput `json:"check_session_input"`
}

// NewChatInput constructor
func NewChatInput() *ChatInput {
	return &ChatInput{}
}
