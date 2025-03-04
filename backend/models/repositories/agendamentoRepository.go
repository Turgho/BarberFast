package repositories

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Estrutura que representa um agendamento
// Contém informações sobre o cliente, serviço, horários e status do agendamento
type Agendamentos struct {
	ID         uint64    `gorm:"primaryKey" json:"id"`
	UsuarioID  string    `gorm:"not null" json:"usuario_id"` // Referência para o ID do usuário (UUID como string)
	ServicoID  uint64    `gorm:"not null" json:"servico_id"`
	DataInicio time.Time `gorm:"not null" json:"data_inicio"`
	DataFim    time.Time `gorm:"not null" json:"data_fim"`
	Status     string    `gorm:"not null" json:"status"` // Exemplo: "confirmado", "cancelado"
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// Relacionamentos
	Usuario Usuarios `gorm:"foreignKey:UsuarioID;constraint:OnDelete:CASCADE" json:"usuario"`
	Servico Servicos `gorm:"foreignKey:ServicoID;constraint:OnDelete:CASCADE" json:"servico"`
}

// Repositório responsável pela manipulação de agendamentos no banco de dados
type AgendamentosRepository struct {
	DB *gorm.DB
}

// Construtor do repositório de agendamentos
func NewAgendamentoRepository(db *gorm.DB) *AgendamentosRepository {
	log.Println("Repositório Agendamentos criado!")
	return &AgendamentosRepository{DB: db}
}

// Cria um novo agendamento no banco de dados
func (repo *AgendamentosRepository) CreateAgendamento(agendamento *Agendamentos) error {
	now := time.Now()

	// Verifica se a data de início é menor ou igual à data atual
	if agendamento.DataInicio.Before(now) || agendamento.DataInicio.Equal(now) {
		log.Println("Erro: a data de início deve ser maior que a data atual")
		return errors.New("a data de início deve ser maior que a data atual")
	}

	// Verifica se a data está dentro do horário comercial permitido
	weekday := agendamento.DataInicio.Weekday()
	hour := agendamento.DataInicio.Hour()

	// Verificação de horário comercial para segunda a sexta das 9h às 17h e sábado das 8h às 13h
	if (weekday >= time.Monday && weekday <= time.Friday && (hour < 9 || hour >= 17)) ||
		(weekday != time.Saturday && (hour < 8 || hour >= 13)) {
		log.Println("Erro: agendamentos permitidos somente de segunda a sexta das 9h às 17h e aos sábados das 8h às 13h")
		return errors.New("agendamentos permitidos somente de segunda a sexta das 9h às 17h e aos sábados das 8h às 13h")
	}

	// Verifica se há conflito com outro agendamento existente
	var existing Agendamentos
	if err := repo.DB.Where("data_inicio < ? AND data_fim > ?", agendamento.DataFim, agendamento.DataInicio).First(&existing).Error; err == nil {
		log.Println("Erro: conflito de horário com outro agendamento")
		return errors.New("conflito de horário com outro agendamento")
	}

	// Cria o agendamento no banco de dados
	if err := repo.DB.Create(agendamento).Error; err != nil {
		log.Printf("Erro ao criar agendamento: %v", err)
		return fmt.Errorf("erro ao criar agendamento: %v", err)
	}

	log.Println("Agendamento criado com sucesso!")
	return nil
}

// Busca um agendamento pelo ID
func (repo *AgendamentosRepository) FindAgendamentoById(agendamentoId string) (*Agendamentos, error) {
	var agendamento Agendamentos

	// Busca o agendamento pelo ID
	if err := repo.DB.Preload("Usuario").Preload("Servico").First(&agendamento, agendamentoId).Error; err != nil {
		log.Printf("Erro ao buscar agendamento: %v", err)
		return nil, fmt.Errorf("erro ao buscar agendamento: %v", err)
	}

	log.Printf("Agendamento encontrado: %v", agendamento)
	return &agendamento, nil
}

// Lista todos os agendamentos conforme um critério de pesquisa
func (repo *AgendamentosRepository) ListAllAgendamentos(nomeUsuario, tipoPesquisa string) ([]Agendamentos, error) {
	var agendamentos []Agendamentos
	now := time.Now()

	// Precarrega as tabelas de Usuários e Serviços
	query := repo.DB.Preload("Usuario").Preload("Servico")

	// Define o tipo de pesquisa
	switch tipoPesquisa {
	case "recente": // Agendamentos mais recentes
		query = query.Where("data_inicio >= ?", now).Order("data_inicio ASC")
	case "distante": // Agendamentos mais distantes
		query = query.Where("data_inicio >= ?", now).Order("data_inicio DESC")
	case "confirmado": // Agendamentos confirmados
		query = query.Where("status = ?", "confirmado")
	case "cancelado": // Agendamentos cancelados
		query = query.Where("status = ?", "cancelado")
	case "nome_cliente":
		query = query.
			Joins("INNER JOIN usuarios ON usuarios.id = agendamentos.usuario_id").
			Where("usuarios.nome LIKE ?", nomeUsuario)
	default:
		log.Println("Erro: tipo de pesquisa inválido")
		return nil, errors.New("tipo de pesquisa inválido")
	}

	// Realiza a query SQL
	if err := query.Find(&agendamentos).Error; err != nil {
		log.Printf("Erro ao buscar agendamentos: %v", err)
		return nil, fmt.Errorf("erro ao buscar agendamentos: %v", err)
	}

	log.Printf("Agendamentos encontrados: %v", len(agendamentos))
	return agendamentos, nil
}

// Deleta um agendamento pelo ID
func (repo *AgendamentosRepository) DeleteAgendamentoById(agendamentoId string) error {
	// Deleta o agendamento pelo ID
	if err := repo.DB.Delete(&Agendamentos{}, agendamentoId).Error; err != nil {
		log.Printf("Erro ao deletar agendamento: %v", err)
		return fmt.Errorf("erro ao deletar agendamento: %v", err)
	}

	log.Printf("Agendamento com ID %s deletado com sucesso!", agendamentoId)
	return nil
}

// Lista todos os agendamentos conforme um critério de pesquisa
func (repo *AgendamentosRepository) ListAgendamentosCliente(ctx *gin.Context, idUsuario, status, ordenacao string) ([]Agendamentos, error) {
	var agendamentos []Agendamentos

	// Mapeia os filtros disponíveis
	filtros := map[string]string{
		"confirmado": "status = 'confirmado'",
		"cancelado":  "status = 'cancelado'",
	}

	// Mapeia a ordenação
	ordenacoes := map[string]string{
		"recente":  "data_inicio ASC",
		"distante": "data_inicio DESC",
	}

	// Inicia a query carregando a relação com `Servico`
	query := repo.DB.WithContext(ctx).Preload("Servico").Where("usuario_id = ?", idUsuario)

	// Aplica filtro de status se necessário
	if filtro, existe := filtros[status]; existe {
		query = query.Where(filtro)
	}

	// Aplica ordenação se necessário
	if ordenacao, existe := ordenacoes[ordenacao]; existe {
		query = query.Order(ordenacao)
	}

	// Executa a busca
	if err := query.Find(&agendamentos).Error; err != nil {
		log.Printf("Erro ao buscar agendamentos: %v", err)
		return nil, fmt.Errorf("erro ao buscar agendamentos: %w", err)
	}

	log.Printf("Agendamentos encontrados: %d", len(agendamentos))
	return agendamentos, nil
}

// Atualiza os agendamentos que passaram do horário
func AtualizarAgendamentosConcluidos(db *gorm.DB) {
	for {
		time.Sleep(5 * time.Minute) // Define a frequência da verificação

		// Atualiza os registros onde a data_fim já passou
		result := db.Model(&Agendamentos{}).
			Where("status <> ? AND data_fim <= ?", "confirmado", time.Now()).
			Update("status", "concluido")

		if result.Error != nil {
			log.Printf("Erro ao atualizar agendamentos: %v", result.Error)
		} else {
			log.Printf("Agendamentos concluídos atualizados: %d", result.RowsAffected)
		}
	}
}
