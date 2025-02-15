package settings

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBConnectionHandler struct {
	DB *mongo.Database
}

func DBConnect(uri, dbName string) (*DBConnectionHandler, error) {
	// Configuração da conexão
	clientOptions := options.Client().ApplyURI(uri)

	// Conexão com MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao mongo_db: %w", err)
	}

	// Verifica se foi conectado
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	fmt.Println("Conectado ao Banco de Dados!")

	// Criar o handler com o banco de dados correto
	handler := &DBConnectionHandler{
		DB: client.Database(dbName),
	}

	return handler, nil
}

func (handler *DBConnectionHandler) Close() {
	if err := handler.DB.Client().Disconnect(context.TODO()); err != nil {
		fmt.Printf("erro ao fechar conexão com DB: %v", err)
	}
}
