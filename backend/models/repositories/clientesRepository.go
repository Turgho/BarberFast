package repositories

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Clientes struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Nome      string    `gorm:"size:255;not null" json:"nome"`
	Email     string    `gorm:"uniqueIndex;size:255;not null" json:"email"`
	Telefone  string    `gorm:"size:14;not null" json:"telefone"`
	Senha     string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type ClientesRepository struct {
	DB *gorm.DB
}

func NewClientesRepository(db *gorm.DB) *ClientesRepository {
	return &ClientesRepository{DB: db}
}

func (repo *ClientesRepository) CreateCliente(cliente *Clientes) error {
	// Cria um novo cliente
	result := repo.DB.Create(&cliente)

	if result.Error != nil {
		return fmt.Errorf("erro ao criar cliente: %v", result.Error)
	}
	return nil
}

func (repo *ClientesRepository) FindClienteById(clienteId string) (*Clientes, error) {
	var cliente Clientes
	result := repo.DB.Where("id = ?", clienteId).Find(&cliente)

	if result.Error != nil {
		return nil, fmt.Errorf("erro ao achar cliente: %v", result.Error)
	}

	return &cliente, nil
}

func (repo *ClientesRepository) ListAllClientes() ([]Clientes, error) {
	var allClientes []Clientes
	result := repo.DB.Find(&allClientes)

	if result.Error != nil {
		return nil, fmt.Errorf("erro ao achar os clientes: %v", result.Error)
	}

	return allClientes, nil
}
