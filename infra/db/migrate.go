package db

import (
	"github.com/schweigert/mmosandbox/domain/entities"
)

// AutoMigrate the database
func AutoMigrate() {
	conn := Connect()
	defer conn.Close()

	conn.AutoMigrate(
		&entities.Account{},
		&entities.Character{},
	)
}
