package main

import (
	"base-api/api/route"
	"base-api/core/app/entity"
	"base-api/core/database"
	"base-api/core/server"
	"base-api/pkg/config"
	"fmt"
	log "github.com/edneyosf/gloged"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB
var e *echo.Echo

func main() {
	loadEnv()
	dbAvailable := initDatabase()

	if dbAvailable {
		startServer()
	}
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		msg := fmt.Sprintf("Error loading .env file: %v", err)

		log.E(msg)
	}
}

func initDatabase() bool {
	var err error
	db, err = database.NewPostgres()

	if err == nil {
		err = db.AutoMigrate(&entity.Base{})
	}

	if err != nil {
		log.E("Shutting down the server: " + err.Error())
	}

	return err == nil
}

func startServer() {
	e = server.New()

	route.SetupRoutes(e, db)

	log.I("Starting API v" + config.Version + " on port " + config.Port)
	err := e.Start(":" + config.Port)

	if err != nil {
		log.E("Shutting down the server: " + err.Error())
	}
}
