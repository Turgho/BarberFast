package handlers

import (
	"fmt"
	"net/http"

	"github.com/Turgho/barberfast/backend/models/repositories"
	"github.com/gin-gonic/gin"
)

var clienteRepo *repositories.ClientesRepository

func InitClientesRepository(repo *repositories.ClientesRepository) {
	clienteRepo = repo
}

func RegistryCliente(ctx *gin.Context) {
	var cliente repositories.Clientes

	// Verifica se JSON tem erro
	if err := ctx.ShouldBindJSON(&cliente); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Cria um cliente no DB
	if err := clienteRepo.CreateCliente(&cliente); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("cliente criado")
	ctx.JSON(http.StatusOK, gin.H{
		"id": cliente.ID,
	})
}

func FindClienteById(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não informado",
		})
		return
	}

	cliente, err := clienteRepo.FindClienteById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		fmt.Println("erro ao achar cliente:", err)
		return
	}

	fmt.Println("cliente encontrado!")
	ctx.JSON(http.StatusOK, cliente)
}

func ListClientes(ctx *gin.Context) {
	allCliente, err := clienteRepo.ListAllClientes()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		fmt.Println("erro ao achar clientes:", err)
		return
	}
	ctx.JSON(http.StatusOK, allCliente)
}
