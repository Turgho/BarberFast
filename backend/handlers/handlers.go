package handlers

import "github.com/Turgho/barberfast/backend/models/repositories"

func InitHandlers(
	clienteRepo *repositories.ClientesRepository,
	servicosRepo *repositories.ServicoRepository,
) {
	InitClientesRepository(clienteRepo)
	InitServicosRepository(servicosRepo)
}
