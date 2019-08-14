package entities

import "math"

// Character model
type Character struct {
	Model

	Name  string `json:"name"`
	Level uint   `json:"level"`
	Exp   uint   `json:"exp"`

	ConsumedStatusPoints uint `json:"consumed_status_point"`
	FreePoints           uint `json:"free_points"`

	StrengthStatus     uint `json:"strength"`
	IntelligenceStatus uint `json:"intelligence"`
	AgilityStatus      uint `json:"agility"`
	VitalityStatus     uint `json:"vitality"`

	MapIndex     uint `json:"map"`
	MapXPosition int  `json:"x_position"`
	MapYPosition int  `json:"y_position"`

	Account   Account `json:"account"`
	AccountID uint    `json:"account_id"`

	MessageBox *MessageBox
}

// DistanceTo other position
func (character *Character) DistanceTo(x int, y int) float64 {
	return math.Sqrt(
		float64(
			(character.MapXPosition-x)*
				(character.MapXPosition-x) +
				(character.MapYPosition-y)*
					(character.MapYPosition-y)))
}

// NewCharacter constructor
func NewCharacter() *Character {
	return &Character{
		Level:              1,
		Exp:                0,
		StrengthStatus:     5,
		IntelligenceStatus: 5,
		AgilityStatus:      5,
		VitalityStatus:     5,
		MapIndex:           0,
		MapXPosition:       0,
		MapYPosition:       0,

		MessageBox: NewMessageBox(),
	}
}
