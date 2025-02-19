package handlers

import (
	"log"
	"net/http"

	"github.com/Turgho/barberfast/backend/models/repositories"
	"github.com/gin-gonic/gin"
)

var servicosRepo *repositories.ServicoRepository

func InitServicosRepository(repo *repositories.ServicoRepository) {
	servicosRepo = repo
}

// Handler de Registro de Serviço
func RegistryServico(ctx *gin.Context) {
	var servico repositories.Servicos

	// Verifica se JSON tem erro
	if err := ctx.ShouldBindJSON(&servico); err != nil {
		log.Printf("Erro ao fazer bind do JSON de serviço: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos",
		})
		return
	}

	// Cria um serviço no DB
	if err := servicosRepo.CreateServico(&servico); err != nil {
		log.Printf("Erro ao criar serviço: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Serviço criado com sucesso! ID: %d", servico.ID)
	ctx.JSON(http.StatusOK, gin.H{
		"id": servico.ID,
	})
}

// Handler de Busca de Serviço por ID
func FindServicoById(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não informado",
		})
		return
	}

	servico, err := servicosRepo.FindServicoById(id)
	if err != nil {
		log.Printf("Erro ao encontrar serviço com ID %s: %v", id, err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Serviço encontrado com sucesso! ID: %d", servico.ID)
	ctx.JSON(http.StatusOK, servico)
}

// Handler de Listagem de Serviços
func ListAllServicos(ctx *gin.Context) {
	servicos, err := servicosRepo.ListAllServicos()
	if err != nil {
		log.Printf("Erro ao listar serviços: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Total de serviços encontrados: %d", len(servicos))
	ctx.JSON(http.StatusOK, servicos)
}

// Handler de Exclusão de Serviço por ID
func DeleteServicoById(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não informado",
		})
		return
	}

	if err := servicosRepo.DeleteServicosById(id); err != nil {
		log.Printf("Erro ao deletar serviço com ID %s: %v", id, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Serviço com ID %s deletado com sucesso", id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Serviço deletado com sucesso!",
	})
}
