package database

import (
	"github.com/jinzhu/gorm"
	// Postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Connect to database
func Connect(conf string) (*gorm.DB, error) {
	return gorm.Open("postgres", conf)
}
