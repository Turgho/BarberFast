package repositories

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Servicos struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Nome          string    `gorm:"size:255;not null" json:"nome"`
	Descricao     string    `gorm:"255;not null" json:"descricao"`
	Preco         float64   `gorm:"not null" json:"preco"`
	DuracaoMinima int       `gorm:"not null" json:"duracao_minutos"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type ServicoRepository struct {
	DB *gorm.DB
}

func NewServicoRepository(db *gorm.DB) *ServicoRepository {
	return &ServicoRepository{DB: db}
}

func (repo *ServicoRepository) CreateServico(servico *Servicos) error {
	result := repo.DB.Create(&servico)

	if result.Error != nil {
		return fmt.Errorf("erro ao criar servico: %v", result.Error)
	}
	return nil
}

func (repo *ServicoRepository) FindServicoById(servicoId string) (*Servicos, error) {
	var servico *Servicos

	result := repo.DB.Where("id = ?", servicoId).Find(&servico)
	if result.Error != nil {
		return nil, fmt.Errorf("erro ao achar servico: %v", result.Error)
	}

	return servico, nil
}

func (repo *ServicoRepository) ListAllServicos() ([]Servicos, error) {
	var servicos []Servicos

	result := repo.DB.Find(&servicos)
	if result.Error != nil {
		return nil, fmt.Errorf("erro ao listar todos os serviços: %v", result.Error)
	}

	return servicos, nil
}
