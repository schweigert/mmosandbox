package inputs

// MoveCharacterInput struct
type MoveCharacterInput struct {
	CharacterID int `json:"character_id"`
	DeltaX      int `json:"delta_x"`
	DeltaY      int `json:"delta_y"`

	CheckSessionInput *CheckSessionInput `json:"check_session_input"`
}

// NewMoveCharacterInput constructor
func NewMoveCharacterInput() *MoveCharacterInput {
	return &MoveCharacterInput{}
}
