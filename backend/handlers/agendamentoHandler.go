package handlers

import (
	"log"
	"net/http"

	"github.com/Turgho/barberfast/backend/models/repositories"
	"github.com/gin-gonic/gin"
)

var agendamentoRepo *repositories.AgendamentosRepository

func InitAgendamentoRepository(repo *repositories.AgendamentosRepository) {
	agendamentoRepo = repo
}

func RegistryAgendamento(ctx *gin.Context) {
	var agendamento repositories.Agendamentos

	// Verifica se JSON tem erro
	if err := ctx.ShouldBindJSON(&agendamento); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Cria um servico no DB
	if err := agendamentoRepo.CreateAgendamento(&agendamento); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("agendamento criado")
	ctx.JSON(http.StatusOK, gin.H{
		"id": agendamento.ID,
	})
}

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
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		log.Fatal("erro ao achar cliente:", err)
		return
	}

	log.Println("agendamento encontrado!")
	ctx.JSON(http.StatusOK, agendamento)
}

func ListAllAgendamentos(ctx *gin.Context) {
	pesquisa := ctx.DefaultQuery("pesquisa", "")
	if pesquisa == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não informado",
		})
		return
	}

	allAgendamentos, err := agendamentoRepo.ListAllAgendamentos(pesquisa)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, allAgendamentos)
}

func DeleteAgendamentoById(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não informado",
		})
		return
	}

	if err := agendamentoRepo.DeleteAgendamentoById(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "agendamento deletado!",
	})
}
