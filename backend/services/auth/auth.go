package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// jwtKey armazena a chave secreta gerada aleatoriamente na inicialização do programa
var jwtKey []byte

// Claims define a estrutura das informações contidas no token JWT
type Claims struct {
	Username             string `json:"username"` // Nome de usuário associado ao token
	Role                 bool   `json:"role"`
	jwt.RegisteredClaims        // Contém campos padrão como expiração
}

// Ele gera uma chave secreta aleatória para assinar os tokens
func init() {
	var err error
	jwtKey, err = generateRandomKey(32) // Gera uma chave de 32 bytes (256 bits)
	if err != nil {
		panic("Erro ao gerar chave secreta: " + err.Error()) // Encerra o programa se houver erro
	}
}

// generateRandomKey gera uma chave secreta aleatória de um determinado tamanho
func generateRandomKey(size int) ([]byte, error) {
	key := make([]byte, size)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return []byte(base64.StdEncoding.EncodeToString(key)), nil
}

// GenerateJWT gera um token JWT para um usuário específico
func GenerateJWT(username string, role bool) (string, error) {
	// Define o tempo de expiração do token (24 horas a partir de agora)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Cria as informações (claims) que serão armazenadas no token
	claims := &Claims{
		Username: username, // Define o nome do usuário no payload do token
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Define a data de expiração
		},
	}

	// Cria um novo token usando o algoritmo HS256 e as informações definidas
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token com a chave secreta e retorna como uma string
	return token.SignedString(jwtKey)
}

// ValidateJWT valida um token JWT e retorna as informações contidas nele
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Faz o parsing do token, validando a assinatura e extraindo os claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if jwtKey == nil {
			return nil, errors.New("chave secreta não inicializada")
		}
		return jwtKey, nil
	})

	// Se houver erro na validação ou o token for inválido, retorna erro
	if err != nil || !token.Valid {
		return nil, err
	}

	// Retorna os claims extraídos do token
	return claims, nil
}

func ExtractClaimsFromToken(ctx *gin.Context) (*Claims, error) {
	// Obtém o cabeçalho Authorization da requisição
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("Token não encontrado")
	}

	// Remove o prefixo "Bearer " do token
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Valida o token JWT
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return nil, errors.New("Token inválido ou expirado")
	}

	return claims, nil
}
