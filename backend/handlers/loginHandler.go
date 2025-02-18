package handlers

import (
	"net/http"

	"github.com/Turgho/barberfast/backend/services/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Struct de Login (username e password)
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Busca o usuário no repositório
	usuario, err := usuariosRepo.GetUsuarioLogin(loginInput.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Usuário não encontrado",
		})
		return
	}

	// Aqui você não precisa gerar o hash novamente, basta comparar a senha fornecida com o hash do banco
	if loginInput.Username != usuario.Nome || !checkPassword(usuario.Senha, loginInput.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário ou senha inválidos"})
		return
	}

	// Supondo que você tenha o papel do usuário em 'usuario.Role'
	token, err := auth.GenerateJWT(loginInput.Username, usuario.IsAdmin) // Passando o role
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o token"})
		return
	}

	// Retorna o token JWT gerado
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Login bem-sucedido",
		"token":   token,
	})
}
