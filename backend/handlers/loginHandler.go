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
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Função para verificar se a senha fornecida corresponde ao hash
func checkPassword(hashedPassword, password string) bool {
	// Compara o hash da senha fornecida com o hash armazenado
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// Handler de Login
func Login(ctx *gin.Context) {
	var loginInput LoginInput

	// Faz o bind das informações enviadas no corpo da requisição (JSON)
	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		log.Printf("Erro ao fazer bind do JSON de login: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Busca o usuário no repositório
	usuario, err := usuariosRepo.GetUsuarioLogin(loginInput.Username)
	if err != nil {
		log.Printf("Usuário não encontrado: %s", loginInput.Username)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	// Verifica se a senha fornecida corresponde ao hash da senha armazenada
	if loginInput.Username != usuario.Nome || !checkPassword(usuario.Senha, loginInput.Password) {
		log.Printf("Falha de login para o usuário: %s", loginInput.Username)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	// Geração do token JWT
	token, err := auth.GenerateJWT(loginInput.Username, usuario.IsAdmin)
	if err != nil {
		log.Printf("Erro ao gerar token para o usuário %s: %v", loginInput.Username, err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o token"})
		return
	}

	log.Printf("Login bem-sucedido para o usuário: %s", loginInput.Username)

	// Retorna o token JWT gerado
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login bem-sucedido",
		"token":   token,
	})
}
