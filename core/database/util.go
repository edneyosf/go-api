package database

import (
	"fmt"
	log "github.com/edneyosf/gloged"
	"gorm.io/gorm"
	"os"
	"time"
)

func loadEnvironmentVariables() bool {
	host = os.Getenv("DB_HOST")
	port = os.Getenv("DB_PORT")
	user = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	name = os.Getenv("DB_NAME")

	return !(host == "" || port == "" || user == "" || password == "" || name == "")
}

func postgresDsn() string {
	const format = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"

	return fmt.Sprintf(format, host, port, user, password, name, sslmode)
}

func mariaDbDsn() string {
	const format = "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	return fmt.Sprintf(format, user, password, host, port, name)
}

func connect(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	// Connection attempt
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(dialector, config)

		if err == nil { break }

		msg := fmt.Sprintf("Failed to connect to database. Retrying in %d seconds... (%d/%d)", delay, i+1, maxRetries)
		log.W(msg)
		time.Sleep(delay * time.Second)
	}

	return db, err
}
