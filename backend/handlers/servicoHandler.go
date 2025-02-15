package handlers

import (
	"fmt"
	"net/http"

	"github.com/Turgho/barberfast/backend/models/repositories"
	"github.com/gin-gonic/gin"
)

var servicosRepo *repositories.ServicoRepository

func InitServicosRepository(repo *repositories.ServicoRepository) {
	servicosRepo = repo
}

func RegistryServico(ctx *gin.Context) {
	var servico repositories.Servicos

	// Verifica se JSON tem erro
	if err := ctx.ShouldBindJSON(&servico); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Cria um servico no DB
	if err := servicosRepo.CreateServico(&servico); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("cliente criado")
	ctx.JSON(http.StatusOK, gin.H{
		"id": servico.ID,
	})
}

func ListAllServicos(ctx *gin.Context) {
	allServico, err := servicosRepo.ListAllServicos()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, allServico)
}
