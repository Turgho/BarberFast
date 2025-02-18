package routes

import (
	"github.com/Turgho/barberfast/backend/handlers"
	"github.com/Turgho/barberfast/backend/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// Rota pública (Login não requer autenticação)
		v1.POST("/login", handlers.Login)

		// Rotas que exigem autenticação (JWTAuthMiddleware)
		v1.Use(middleware.JWTAuthMiddleware()) // Middleware aplicado a todas as rotas abaixo

		// Grupo admin com permissão extra
		admin := v1.Group("/admin")
		{
			// Middleware adicional para verificar se é admin
			admin.Use(middleware.AdminMiddleware())

			// Rotas administrativas
			admin.GET("/ping", handlers.PingHandler)

			admin.GET("/clientes", handlers.ListAllUsuarios)
			admin.GET("/cliente", handlers.FindUsuarioById)
			admin.DELETE("/cliente", handlers.DeleteUsuario)

			admin.POST("/servicos", handlers.RegistryServico)
			admin.GET("/servicos", handlers.ListAllServicos)
			admin.GET("/servico", handlers.FindServicoById)
			admin.DELETE("/servico", handlers.DeleteServicoById)

			admin.GET("/agendamento", handlers.FindAgendamentoById)
			admin.GET("/agendamentos", handlers.ListAllAgendamentos)
			admin.DELETE("/agendamento", handlers.DeleteAgendamentoById)
		}

		// Rotas para o usuário comum
		usuario := v1.Group("/usuario")
		{
			usuario.POST("/clientes", handlers.RegistryUsuario)
			usuario.POST("/agendamentos", handlers.RegistryAgendamento)
		}
	}
}
