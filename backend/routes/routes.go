package routes

import (
	_ "github.com/Turgho/barberfast/backend/docs"

	"github.com/Turgho/barberfast/backend/handlers"
	"github.com/Turgho/barberfast/backend/middleware"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	// Rota para documentação Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Adiciona o middleware de CORS
	r.Use(middleware.CORSMiddleware())

	v1 := r.Group("/v1")
	{
		// Rota pública (Login/Cadastro não requer autenticação)
		v1.POST("/login", handlers.Login)
		v1.POST("/cadastro", handlers.RegistryUsuario)

		// Middleware de autenticação aplicado a todas as rotas abaixo
		v1.Use(middleware.JWTAuthMiddleware())

		// Grupo admin (Permissões extras)
		admin := v1.Group("/admin")
		admin.Use(middleware.AdminMiddleware()) // Aplica o middleware só para admin
		{
			// Rotas administrativas
			admin.GET("/ping", handlers.PingHandler)

			// Rotas de clientes
			admin.GET("/clientes", handlers.ListAllUsuarios)
			admin.GET("/cliente", handlers.FindUsuarioById)
			admin.DELETE("/cliente", handlers.DeleteUsuario)

			// Rotas de serviços
			admin.POST("/servicos", handlers.RegistryServico)
			admin.GET("/servicos", handlers.ListAllServicos)
			admin.GET("/servico", handlers.FindServicoById)
			admin.DELETE("/servico", handlers.DeleteServicoById)

			// Rotas de agendamentos
			admin.GET("/agendamento", handlers.FindAgendamentoById)
			admin.GET("/agendamentos", handlers.ListAllAgendamentos)
			admin.DELETE("/agendamento", handlers.DeleteAgendamentoById)
		}

		// Rotas para usuários comuns
		usuario := v1.Group("/usuario")
		{ // Ajuste para singular
			usuario.POST("/agendamento", handlers.RegistryAgendamento) // Ajuste para singular
		}
	}
}
