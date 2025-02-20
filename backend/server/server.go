package server

import (
	"log"

	"github.com/Turgho/barberfast/backend/handlers"
	"github.com/Turgho/barberfast/backend/migration"
	"github.com/Turgho/barberfast/backend/models/repositories"
	"github.com/Turgho/barberfast/backend/models/settings"
	"github.com/Turgho/barberfast/backend/routes"
	"github.com/Turgho/barberfast/backend/services/rabbitmq"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	// Inicia o log da API
	settings.SetupLogging()

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

	// Inicia o RabbitMQ
	if err := rabbitmq.SendMessageToQueue("TESTE MENSAGEM!"); err != nil {
		log.Fatal("erro ao iniciar RabbitMQ")
	}

	// Consume mensagens de forma assíncrona em uma goroutine
	// go rabbitmq.ConsumeMessages()

	// Inicia o router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Inicia o servidor na porta :8080
	if err := router.Run(":5050"); err != nil {
		log.Fatal("erro ao iniciar ao servidor:", err)
		return
	}
}
