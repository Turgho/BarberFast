package handlers

import (
	"log"

	"github.com/Turgho/barberfast/backend/models/repositories"
)

func InitHandlers(
	usuarioRepo *repositories.UsuariosRepository,
	servicosRepo *repositories.ServicoRepository,
	agendamentoRepo *repositories.AgendamentosRepository,
) {
	InitUsuariosRepository(usuarioRepo)
	InitServicosRepository(servicosRepo)
	InitAgendamentoRepository(agendamentoRepo)

	log.Println("Handlers carregados!")
}
