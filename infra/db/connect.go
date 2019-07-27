package db

import (
	"github.com/jinzhu/gorm"
	"github.com/schweigert/mmosandbox/config"
	"github.com/schweigert/mmosandbox/lib/database"
	"github.com/schweigert/mmosandbox/lib/dont"
)

// Connect to PG
func Connect() *gorm.DB {
	conn, err := database.Connect(config.Db().Addr())
	dont.Panic(err)

	return conn
}
