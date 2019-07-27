package db

import (
	"github.com/jinzhu/gorm"
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/domain/entities"
	"github.com/schweigert/mmosandbox/lib/database"
	"github.com/schweigert/mmosandbox/lib/dont"
)

// Connect to PG
func Connect() *gorm.DB {
	conn, err := database.Connect(config.Db().Addr())
	dont.Panic(err)

	return conn
}

// Migrate the database
func Migrate() {
	conn := Connect()
	defer conn.Close()

	conn.AutoMigrate(
		&entities.Account{},
		&entities.Character{},
	)
}

// Clear the database
func Clear() {
	conn := Connect()
	defer conn.Close()

	conn.DropTableIfExists(
		&entities.Account{},
		&entities.Character{},
	)
}
