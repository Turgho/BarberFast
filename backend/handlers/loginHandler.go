package handlers

import (
	"log"
	"net/http"

	"github.com/Turgho/barberfast/backend/services/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Estrutura de Login (username e password)
type LoginInput struct {
	Email string `json:"email" binding:"required"`
	Senha string `json:"senha" binding:"required"`
}

// Função para verificar se a senha fornecida corresponde ao hash
func checkPassword(hashedPassword, senha string) bool {
	// Compara o hash da senha fornecida com o hash armazenado
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(senha))
	return err == nil
}

// Handler de Login

// Login autentica um usuário e retorna um token JWT.
//
// @Summary      Autenticação de usuário
// @Description  Autentica um usuário e retorna um token de acesso
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        login  body  map[string]string  true  "Credenciais de login"
// @Success      200    {object}  map[string]interface{}
// @Failure      401    {object}  map[string]string "Usuário ou senha inválidos"
// @Router       /v1/login [post]
func Login(ctx *gin.Context) {
	var loginInput LoginInput

	// Faz o bind das informações enviadas no corpo da requisição (JSON)
	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		log.Printf("Erro ao fazer bind do JSON de login: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Busca o usuário no repositório
	usuario, err := usuariosRepo.GetUsuarioLogin(loginInput.Email)
	if err != nil {
		log.Printf("Usuário não encontrado: %s", loginInput.Email)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	// Verifica se a senha fornecida corresponde ao hash da senha armazenada
	if loginInput.Email != usuario.Email || !checkPassword(usuario.Senha, loginInput.Senha) {
		log.Printf("Falha de login para o usuário: %s", loginInput.Email)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	// Geração do token JWT
	token, err := auth.GenerateJWT(loginInput.Email, usuario.IsAdmin)
	if err != nil {
		log.Printf("Erro ao gerar token para o usuário %s: %v", usuario.Nome, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o token"})
		return
	}

	log.Printf("Login bem-sucedido para o usuário: %s", usuario.Nome)

	// Retorna o token JWT gerado
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Login bem-sucedido",
		"user_id":  usuario.ID,
		"username": usuario.Nome,
		"token":    token,
	})
}
