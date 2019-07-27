package entities

import "github.com/jinzhu/gorm"

// Account model
type Account struct {
	*gorm.Model

	Username       string `gorm:"unique_index" json:"username"`
	SecurePassword string `json:"secure_password"`
	Email          string `json:"email"`

	Characters []Character `json:"characters"`
}

// NewAccount constructor
func NewAccount() *Account {
	return &Account{}
}

// Character model
type Character struct {
	*gorm.Model

	Name  string `gorm:"unique_index" json:"name"`
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
}

// NewCharacter constructor
func NewCharacter() *Character {
	return &Character{}
}
