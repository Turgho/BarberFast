// @title           BarberFast API
// @version         1.0
// @description     API para gerenciamento de barbearia
// @host           localhost:5050
// @BasePath       /api/v1
package main

import (
	"log"

	"github.com/Turgho/barberfast/backend/server"
)

func main() {
	server.InitServer()

	log.Println("API INICIADA!")
}
