package entities

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
	}
}
