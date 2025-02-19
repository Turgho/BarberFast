package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Turgho/barberfast/backend/models/repositories"
	"github.com/Turgho/barberfast/backend/services/security"
	"github.com/gin-gonic/gin"
)

var usuariosRepo *repositories.UsuariosRepository

func InitUsuariosRepository(repo *repositories.UsuariosRepository) {
	usuariosRepo = repo
}

func RegistryUsuario(ctx *gin.Context) {
	var usuario repositories.Usuarios

	// Verifica se JSON tem erro
	if err := ctx.ShouldBindJSON(&usuario); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Criptografa a senha do usuário antes de salvar

	hashedSenha, err := security.HashPassword(usuario.Senha)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	// Atribui a senha criptografada ao usuário
	usuario.Senha = hashedSenha

	// Cria um cliente no DB
	if err := usuariosRepo.CreateUsuario(&usuario); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("cliente criado")
	ctx.JSON(http.StatusCreated, gin.H{
		"id": usuario.ID,
	})
}

func FindUsuarioById(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não informado",
		})
		return
	}

	cliente, err := usuariosRepo.FindUsuarioById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		log.Fatal("erro ao achar usuario:", err)
		return
	}

	fmt.Println("cliente encontrado!")
	ctx.JSON(http.StatusOK, cliente)
}

func ListAllUsuarios(ctx *gin.Context) {
	allUsuarios, err := usuariosRepo.ListAllUsuarios()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		log.Fatal("erro ao achar usuários:", err)
		return
	}
	ctx.JSON(http.StatusOK, allUsuarios)
}

func DeleteUsuario(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não informado",
		})
		return
	}

	if err := usuariosRepo.DeleteUsuarioById(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		log.Fatal("erro ao deletar usuário:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "usuário deletado!",
	})
}
