package db

import (
	"github.com/schweigert/mmosandbox/domain/entities"
)

// Clear the database
func Clear() {
	conn := Connect()
	defer conn.Close()

	conn.DropTableIfExists(
		&entities.Account{},
		&entities.Character{},
	)

	AutoMigrate()
}
