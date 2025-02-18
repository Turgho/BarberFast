package middleware

import (
	"net/http"

	"github.com/Turgho/barberfast/backend/services/auth"
	"github.com/gin-gonic/gin"
)

// Ele verifica a existência e validade do token antes de permitir o acesso às rotas protegidas
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := auth.ExtractClaimsFromToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		// Armazena o nome de usuário extraído do token no contexto da requisição
		ctx.Set("username", claims.Username)

		ctx.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, err := auth.ExtractClaimsFromToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		// Verifica se o usuário tem permissão de admin
		if !claims.Role {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Acesso negado"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
