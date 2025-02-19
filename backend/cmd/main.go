package main

import (
	"log"

	"github.com/Turgho/barberfast/backend/server"
)

func main() {
	server.InitServer()

	log.Println("API INICIADA!")
}
