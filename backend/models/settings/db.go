package settings

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConnectionHandler é um tipo que encapsula a conexão com o banco de dados, permitindo a manipulação do objeto GORM.
type DBConnectionHandler struct {
	DB *gorm.DB
}

// DBConnect é responsável por estabelecer a conexão com o banco de dados, utilizando as credenciais do arquivo .env.
func DBConnect() (*DBConnectionHandler, error) {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	// Recupera as credenciais de conexão do banco de dados das variáveis de ambiente
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),     // Usuário do banco de dados
		os.Getenv("DB_PASSWORD"), // Senha do banco de dados
		os.Getenv("DB_HOST"),     // Endereço do host do banco de dados
		os.Getenv("DB_NAME"),     // Nome do banco de dados
	)

	// Tenta abrir a conexão com o banco de dados utilizando o driver MySQL
	handler, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Erro ao abrir conexão com o banco de dados: %v", err)
		return nil, errors.New("erro ao abrir conexão com banco de dados")
	}

	// Obtém a instância do banco de dados subjacente
	db, err := handler.DB()
	if err != nil {
		log.Printf("Erro ao acessar o banco de dados: %v", err)
		return nil, errors.New("erro ao acessar banco de dados")
	}

	log.Println("Conectado ao Banco de Dados!")

	// Verifica se é possível acessar o banco de dados
	if err := db.Ping(); err != nil {
		log.Printf("Erro ao acessar o banco de dados: %v", err)
		return nil, errors.New("erro ao acessar banco de dados")
	}

	return &DBConnectionHandler{DB: handler}, nil
}

// Close encerra a conexão com o banco de dados de maneira adequada.
func (handler *DBConnectionHandler) Close() {
	// Obtém a instância do banco de dados subjacente para poder fechar a conexão
	db, err := handler.DB.DB()
	if err != nil {
		log.Printf("Erro ao acessar o banco de dados para fechar a conexão: %v", err)
		return
	}

	// Tenta fechar a conexão com o banco de dados
	if err := db.Close(); err != nil {
		log.Printf("Erro ao fechar conexão com banco de dados: %v", err)
	} else {
		log.Println("Conexão com o banco de dados fechada com sucesso.")
	}
}
