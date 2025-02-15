package routes

import (
	"github.com/Turgho/barberfast/backend/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.PingHandler)

	r.GET("/clientes", handlers.ListClientes)
	r.POST("/clientes", handlers.RegistryCliente)
	r.GET("/cliente", handlers.FindClienteById)

	r.POST("/servicos", handlers.RegistryServico)
	r.GET("/servicos", handlers.ListAllServicos)
}
