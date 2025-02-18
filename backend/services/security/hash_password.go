package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Função para criptografar a senha
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("senha não pode ser vazia")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
