package server

import (
	"fmt"

	"github.com/Turgho/barberfast/backend/handlers"
	"github.com/Turgho/barberfast/backend/migration"
	"github.com/Turgho/barberfast/backend/models/repositories"
	"github.com/Turgho/barberfast/backend/models/settings"
	"github.com/Turgho/barberfast/backend/routes"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	// Conecta ao DB
	dbHandler, err := settings.DBConnect()
	if err != nil {
		fmt.Printf("erro ao conectar ao DB: %v", err)
	}
	defer dbHandler.Close()

	// Inicia as migrações
	migration.InitMigrations(dbHandler.DB)

	// Inicia Repositories
	clienteRepo := repositories.NewClientesRepository(dbHandler.DB)
	servicoRepo := repositories.NewServicoRepository(dbHandler.DB)

	// Inicia Handlers
	handlers.InitHandlers(
		clienteRepo,
		servicoRepo,
	)

	// Inicia o router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Inicia o servidor na porta :8080
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("erro ao iniciar ao servidor: %v", err)
	}
}
