package settings

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConnectionHandler struct {
	DB *gorm.DB
}

func DBConnect() (*DBConnectionHandler, error) {
	// Carregar o arquivo .env
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Erro ao carregar o arquivo .env: ", err)
	}

	// Pegando as credenciais do .env ou configurando manualmente
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),     // Usuário
		os.Getenv("DB_PASSWORD"), // Senha
		os.Getenv("DB_HOST"),     // Host
		os.Getenv("DB_NAME"),     // Nome do banco de dados
	)
	// Abrindo conexão
	handler, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com banco de dados: %v", err)
	}

	// Conectando
	db, err := handler.DB()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)
	}

	fmt.Println("Conectado ao Banco de Dados!")

	// Acessando
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao acessar banco de dados: %v", err)
	}

	return &DBConnectionHandler{DB: handler}, nil
}

func (handler *DBConnectionHandler) Close() {
	db, err := handler.DB.DB()
	if err != nil {
		fmt.Printf("erro ao fechar e acessar banco de dados: %v", err)
	}

	if err := db.Close(); err != nil {
		fmt.Printf("erro ao fechar conexão com banco de dados: %v", err)
	}
}
