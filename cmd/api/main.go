package main

import (
	"base-api/api/routes"
	"base-api/internal/app/controller"
	"base-api/internal/app/entity"
	"base-api/internal/app/repository"
	"base-api/internal/app/usecase"
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"

	//"gorm.io/gorm/logger"
	"log"
)

func main() {
	// Inicializar configuração e logger
	//config.Load()
	//logger.Init()

	// Obter variáveis de ambiente para o banco de dados
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Database environment variables not set")
	}

	// Configurar conexão com o banco de dados PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	var db *gorm.DB
	var err error

	// Tentativa de conexão com retry
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("Failed to connect to database. Retrying in 5 seconds... (%d/%d)", i+1, maxRetries)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		log.Fatalf("Could not connect to the database after %d attempts: %v", maxRetries, err)
	}

	// Migrar a estrutura do usuário
	db.AutoMigrate(&entity.Base{})

	// Configurar dependências
	userRepo := repository.NewRepository(db)
	userUsecase := usecase.NewUsecase(userRepo)
	userController := controller.NewController(userUsecase)

	// Criar instância do Echo
	e := echo.New()

	// Configurar rotas
	routes.SetupRoutes(e, userController)

	// Iniciar servidor
	log.Println("Starting API server on port 5000...")
	if err := e.Start(":5000"); err != nil {
		log.Fatal("Shutting down the server: " + err.Error())
	}
}
