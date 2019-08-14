package entities

// Message recevied from one point in map
type Message struct {
	CharacterID   uint
	CharacterName string
	MapXPosition  int
	MapYPosition  int
	Body          string
}

// NewMessage constructor
func NewMessage() *Message {
	return &Message{}
}

// NewMessageFromCharacter constructor
func NewMessageFromCharacter(character *Character, body string) *Message {
	return &Message{
		CharacterID:   character.ID,
		CharacterName: character.Name,
		MapXPosition:  character.MapXPosition,
		MapYPosition:  character.MapYPosition,
		Body:          body,
	}
}
