package repositories

import (
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

// Servicos representa a estrutura de um serviço oferecido na plataforma.
type Servicos struct {
	ID            uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Nome          string    `gorm:"size:255;not null" json:"nome"`
	Descricao     string    `gorm:"size:255;not null" json:"descricao"`
	Preco         float64   `gorm:"not null" json:"preco"`
	DuracaoMinima int       `gorm:"not null" json:"duracao_minutos"`
	Status        string    `gorm:"not null" json:"status"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// ServicoRepository encapsula a conexão com o banco de dados e fornece métodos para interagir com a tabela de serviços.
type ServicoRepository struct {
	DB *gorm.DB
}

// NewServicoRepository cria e retorna uma nova instância de ServicoRepository.
func NewServicoRepository(db *gorm.DB) *ServicoRepository {
	log.Println("Repositório Serviços criado!")
	return &ServicoRepository{DB: db}
}

// CreateServico cria um novo serviço no banco de dados.
func (repo *ServicoRepository) CreateServico(servico *Servicos) error {
	// Cria o serviço no banco de dados
	result := repo.DB.Create(&servico)

	// Se houve erro ao criar o serviço, retorna um erro detalhado
	if result.Error != nil {
		log.Printf("Erro ao criar serviço: %v", result.Error)
		return errors.New("erro ao criar serviço")
	}
	log.Println("Serviço criado com sucesso!")
	return nil
}

// FindServicoById busca um serviço no banco de dados pelo ID.
func (repo *ServicoRepository) FindServicoById(servicoId string) (*Servicos, error) {
	var servico Servicos

	// Busca o serviço pelo ID
	result := repo.DB.Where("id = ?", servicoId).First(&servico)
	if result.Error != nil {
		log.Printf("Erro ao achar serviço por ID: %v", result.Error)
		return nil, errors.New("erro ao achar serviço")
	}

	log.Printf("Serviço encontrado: %v", servico)
	return &servico, nil
}

func (repo *ServicoRepository) ListServicosDisponiveis() ([]Servicos, error) {
	var servicos []Servicos

	result := repo.DB.
		Where("status = ?", "disponivel").
		Find(&servicos)

	if result.Error != nil {
		log.Printf("Erro ao listar todos os serviços disponiveis: %v", result.Error)
		return nil, errors.New("erro ao listar todos os serviços disponiveis")
	}

	log.Printf("Serviços disponiveis encontrados: %v", len(servicos))
	return servicos, nil
}

// ListAllServicos lista todos os serviços disponíveis no banco de dados.
func (repo *ServicoRepository) ListAllServicos() ([]Servicos, error) {
	var servicos []Servicos

	// Lista todos os serviços
	result := repo.DB.Find(&servicos)
	if result.Error != nil {
		log.Printf("Erro ao listar todos os serviços: %v", result.Error)
		return nil, errors.New("erro ao listar todos os serviços")
	}

	log.Printf("Serviços encontrados: %v", len(servicos))
	return servicos, nil
}

// DeleteServicosById deleta um serviço do banco de dados pelo ID.
func (repo *ServicoRepository) DeleteServicosById(servicoId string) error {
	var servico Servicos

	// Deleta o serviço pelo ID
	result := repo.DB.Where("id = ?", servicoId).Delete(&servico)
	if result.Error != nil {
		log.Printf("Erro ao deletar serviço por ID: %v", result.Error)
		return errors.New("erro ao deletar serviço")
	}

	log.Printf("Serviço com ID %s deletado com sucesso!", servicoId)
	return nil
}
