package server

import (
	"log"

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
		log.Fatal("erro ao conectar ao DB:", err)
		return
	}
	defer dbHandler.Close()

	// Inicia as migrações
	migration.InitMigrations(dbHandler.DB)

	// Inicia Repositories
	clienteRepo := repositories.NewUsuariosRepository(dbHandler.DB)
	servicoRepo := repositories.NewServicoRepository(dbHandler.DB)
	agendamentoRepo := repositories.NewAgendamentoRepository(dbHandler.DB)

	// Inicia Handlers
	handlers.InitHandlers(
		clienteRepo,
		servicoRepo,
		agendamentoRepo,
	)

	// Inicia o router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Inicia o servidor na porta :8080
	if err := router.Run(":5050"); err != nil {
		log.Fatal("erro ao iniciar ao servidor:", err)
		return
	}
}
