package server

import (
	"fmt"
	"os"
	"strings"

	"github.com/Turgho/barberfast/backend/models/settings"
	"github.com/Turgho/barberfast/backend/routes"
	"github.com/gin-gonic/gin"
)

func checkEnvVariables(vars ...string) bool {
	for _, v := range vars {
		if v == "" {
			return false
		}
	}
	return true
}

func InitServer() {
	// Carrega e verifica as variável .ENV
	clusterPassword := os.Getenv("CLUSTER_PASSWORD")
	uri := os.Getenv("URI_STRING")
	databaseName := os.Getenv("DATABASE_NAME")

	if !checkEnvVariables(clusterPassword, uri, databaseName) {
		fmt.Printf("erro ao carregar .ENV")
		return
	}

	uri_string := strings.ReplaceAll(uri, "${CLUSTER_PASSWORD}", databaseName)

	// Conecta ao DB
	dbHandler, err := settings.DBConnect(uri_string, databaseName)
	if err != nil {
		fmt.Printf("erro ao conectar ao DB: %v", err)
	}
	defer dbHandler.Close()

	// Inicia o router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Inicia o servidor na porta :8080
	if err := router.Run(":8080"); err != nil {
		fmt.Printf("erro ao iniciar ao servidor: %v", err)
	}
}
