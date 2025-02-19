package handlers

import (
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

// Handler para registrar um novo usuário
func RegistryUsuario(ctx *gin.Context) {
	var usuario repositories.Usuarios

	// Verifica se JSON tem erro
	if err := ctx.ShouldBindJSON(&usuario); err != nil {
		log.Printf("Erro ao fazer bind de dados de usuário: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Dados inválidos",
		})
		return
	}

	// Criptografa a senha do usuário antes de salvar
	hashedSenha, err := security.HashPassword(usuario.Senha)
	if err != nil {
		log.Printf("Erro ao criptografar senha do usuário: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao processar a senha",
		})
		return
	}

	// Atribui a senha criptografada ao usuário
	usuario.Senha = hashedSenha

	// Cria um cliente no DB
	if err := usuariosRepo.CreateUsuario(&usuario); err != nil {
		log.Printf("Erro ao criar usuário no banco de dados: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao criar usuário",
		})
		return
	}

	log.Printf("Usuário criado com sucesso! ID: %d", usuario.ID)
	ctx.JSON(http.StatusCreated, gin.H{
		"id": usuario.ID,
	})
}

// Handler para buscar um usuário pelo ID
func FindUsuarioById(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não informado",
		})
		return
	}

	usuario, err := usuariosRepo.FindUsuarioById(id)
	if err != nil {
		log.Printf("Erro ao buscar usuário com ID %s: %v", id, err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	log.Printf("Usuário encontrado! ID: %d", usuario.ID)
	ctx.JSON(http.StatusOK, usuario)
}

// Handler para listar todos os usuários
func ListAllUsuarios(ctx *gin.Context) {
	allUsuarios, err := usuariosRepo.ListAllUsuarios()
	if err != nil {
		log.Printf("Erro ao listar usuários: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao listar usuários",
		})
		return
	}

	log.Printf("Total de usuários encontrados: %d", len(allUsuarios))
	ctx.JSON(http.StatusOK, allUsuarios)
}

// Handler para deletar um usuário pelo ID
func DeleteUsuario(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID não informado",
		})
		return
	}

	if err := usuariosRepo.DeleteUsuarioById(id); err != nil {
		log.Printf("Erro ao deletar usuário com ID %s: %v", id, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao deletar usuário",
		})
		return
	}

	log.Printf("Usuário com ID %s deletado com sucesso", id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Usuário deletado com sucesso!",
	})
}
