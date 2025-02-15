package routes

import (
	"github.com/Turgho/barberfast/backend/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/ping", handlers.PingHandler)
}
