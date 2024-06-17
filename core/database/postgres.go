package database

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres() (*gorm.DB, error) {
	if !checkEnvironmentVariables() { return nil, errors.New(noEnvironmentVariablesError) }

	dsn := postgresDsn()
	dialector := postgres.Open(dsn)
	config := &gorm.Config{}

	db, err := connect(dialector, config)

	return db, err
}
