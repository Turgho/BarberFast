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

func getIDFromQuery(ctx *gin.Context) (string, error) {
	id := ctx.DefaultQuery("id", "") // Pega o parâmetro 'id' da query string
	if id == "" {
		return "", fmt.Errorf("ID não encontrado na query string")
	}

	return id, nil
}

// Handler para registrar um novo usuário
func RegistryUsuario(ctx *gin.Context) {
	var usuario repositories.Usuarios

	// Verifica se JSON tem erro
	if err := ctx.ShouldBindJSON(&usuario); err != nil {
		log.Printf("Erro ao fazer bind de dados de usuário: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Criptografa a senha do usuário antes de salvar
	hashedSenha, err := security.HashPassword(usuario.Senha)
	if err != nil {
		log.Printf("Erro ao criptografar senha do usuário: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Atribui a senha criptografada ao usuário
	usuario.Senha = hashedSenha

	// Cria um cliente no DB
	if err := usuariosRepo.CreateUsuario(&usuario); err != nil {
		log.Printf("Erro ao criar usuário no banco de dados: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Printf("Usuário criado com sucesso! ID: %s", usuario.ID)
	ctx.JSON(http.StatusCreated, gin.H{
		"id": usuario.ID,
	})
}

// Handler para buscar um usuário pelo ID

// FindUsuarioById retorna um cliente específico.
//
// @Summary      Buscar cliente por ID
// @Description  Retorna os dados de um cliente pelo seu ID
// @Tags         Clientes
// @Produce      json
// @Param        id  query  string  true  "ID do Cliente"
// @Success      200  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]string "Cliente não encontrado"
// @Router       /v1/admin/cliente [get]
func FindUsuarioById(ctx *gin.Context) {
	id, err := getIDFromQuery(ctx)
	if err != nil {
		log.Printf("Erro ao encontrar parâmetro de ID: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	usuario, err := usuariosRepo.FindUsuarioById(id)
	if err != nil {
		// Aqui a gente pode verificar se o erro é de "não encontrado" ou outro tipo
		log.Printf("Erro ao buscar usuário com ID %s: %v", id, err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Printf("Usuário encontrado! ID: %s", usuario.ID)
	ctx.JSON(http.StatusFound, usuario)
}

// Handler para listar todos os usuários

// ListAllUsuarios retorna todos os clientes cadastrados.
//
// @Summary      Lista todos os clientes
// @Description  Obtém uma lista de todos os clientes registrados no sistema
// @Tags         Clientes
// @Produce      json
// @Success      200    {array}  map[string]interface{}
// @Router       /v1/admin/clientes [get]
func ListAllUsuarios(ctx *gin.Context) {
	allUsuarios, err := usuariosRepo.ListAllUsuarios()
	if err != nil {
		log.Printf("Erro ao listar usuários: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Printf("Total de usuários encontrados: %d", len(allUsuarios))
	ctx.JSON(http.StatusFound, allUsuarios)
}

// Handler para deletar um usuário pelo ID
func DeleteUsuario(ctx *gin.Context) {
	id, err := getIDFromQuery(ctx)
	if err != nil {
		log.Printf("Erro ao encontrar parametro de ID: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	if err := usuariosRepo.DeleteUsuarioById(id); err != nil {
		log.Printf("Erro ao deletar usuário com ID %s: %v", id, err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("Usuário com ID %s deletado com sucesso", id)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Usuário deletado com sucesso!",
	})
}
