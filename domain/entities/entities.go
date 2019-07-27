package entities

import "github.com/jinzhu/gorm"

// Account model
type Account struct {
	*gorm.Model

	Username       string `gorm:"unique_index"`
	SecurePassword string
	Email          string

	Characters []Character
}

// Character model
type Character struct {
	*gorm.Model

	Name  string `gorm:"unique_index"`
	Level uint
	Exp   uint

	ConsumedStatusPoints uint
	FreePoints           uint

	StrengthStatus     uint
	IntelligenceStatus uint
	AgilityStatus      uint
	VitalityStatus     uint

	MapIndex     uint
	MapXPosition int
	MapYPosition int

	Account   Account
	AccountID uint
}
