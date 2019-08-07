package inputs

// SpawnCharacterInput struct
type SpawnCharacterInput struct {
	CharacterID       int                `json:"character_id"`
	CheckSessionInput *CheckSessionInput `json:"check_session_input"`
}

// NewSpawnCharacterInput constructor
func NewSpawnCharacterInput() *SpawnCharacterInput {
	return &SpawnCharacterInput{}
}
