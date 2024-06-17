package database

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMariaDB() (*gorm.DB, error) {
	if !checkEnvironmentVariables() { return nil, errors.New(noEnvironmentVariablesError) }

	dsn := mariaDbDsn()
	dialector := mysql.Open(dsn)
	config := &gorm.Config{}

	db, err := connect(dialector, config)

	return db, err
}
