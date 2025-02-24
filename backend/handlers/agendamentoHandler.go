package handlers

import (
	"log"
	"net/http"

	"github.com/Turgho/barberfast/backend/models/repositories"
	"github.com/gin-gonic/gin"
)

var agendamentoRepo *repositories.AgendamentosRepository

// InitAgendamentoRepository inicializa o repositório de agendamentos
func InitAgendamentoRepository(repo *repositories.AgendamentosRepository) {
	agendamentoRepo = repo
}

// RegistryAgendamento cria um novo agendamento

// RegistryAgendamento cria um novo agendamento.
//
// @Summary      Criar agendamento
// @Description  Permite que o usuário marque um horário no sistema
// @Tags         Agendamentos
// @Accept       json
// @Produce      json
// @Param        agendamento  body  map[string]interface{}  true  "Dados do agendamento"
// @Success      201  {object}  map[string]string  "Agendamento criado com sucesso"
// @Failure      400  {object}  map[string]string  "Erro ao criar agendamento"
// @Router       /v1/usuario/agendamento [post]
func RegistryAgendamento(ctx *gin.Context) {
	var agendamento repositories.Agendamentos

	// Verifica se o JSON está correto
	if err := ctx.ShouldBindJSON(&agendamento); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Tenta criar o agendamento no banco de dados
	if err := agendamentoRepo.CreateAgendamento(&agendamento); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Agendamento criado com sucesso! ID: %d", agendamento.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"id": agendamento.ID,
	})
}

// FindAgendamentoById busca um agendamento pelo ID
func FindAgendamentoById(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não informado",
		})
		return
	}

	agendamento, err := agendamentoRepo.FindAgendamentoById(id)
	if err != nil {
		log.Printf("Erro ao buscar agendamento com ID %s: %v", id, err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Agendamento encontrado! ID: %d", agendamento.ID)
	ctx.JSON(http.StatusOK, agendamento)
}

// ListAllAgendamentos lista todos os agendamentos com base no critério de pesquisa
func ListAllAgendamentos(ctx *gin.Context) {
	pesquisa := ctx.DefaultQuery("pesquisa", "")

	if pesquisa == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Tipo de pesquisa ou usuário não informado",
		})
		return
	}

	var usuario repositories.Usuarios

	// Verifica se o JSON está correto
	if err := ctx.ShouldBindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	allAgendamentos, err := agendamentoRepo.ListAllAgendamentos(usuario.Nome, pesquisa)
	if err != nil {
		log.Printf("Erro ao listar agendamentos: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Encontrados %d agendamentos", len(allAgendamentos))
	ctx.JSON(http.StatusOK, allAgendamentos)
}

// DeleteAgendamentoById deleta um agendamento pelo ID
func DeleteAgendamentoById(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não informado",
		})
		return
	}

	if err := agendamentoRepo.DeleteAgendamentoById(id); err != nil {
		log.Printf("Erro ao deletar agendamento com ID %s: %v", id, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Agendamento com ID %s deletado com sucesso!", id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Agendamento deletado com sucesso!",
	})
}
