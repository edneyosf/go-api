package main

import (
	"base-api/api/route"
	"base-api/core/app/entity"
	"base-api/core/database"
	"base-api/core/server"
	"base-api/pkg/config"
	log "github.com/edneyosf/gloged"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var db *gorm.DB
var e *echo.Echo

func main() {
	dbAvailable := initDatabase()

	if dbAvailable {
		startServer()
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
