package repositories

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

// Definição do modelo de dados para a tabela 'usuarios'
// 'Usuarios' mapeia as colunas do banco de dados com as propriedades do struct.
type Usuarios struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`         // Chave primária e autoincremento
	Nome      string    `gorm:"size:255;not null" json:"nome"`              // Nome do usuário
	Email     string    `gorm:"uniqueIndex;size:255;not null" json:"email"` // Email do usuário, único no banco
	Telefone  string    `gorm:"size:14;not null" json:"telefone"`           // Telefone do usuário
	Senha     string    `gorm:"not null" json:"senha"`                      // Senha do usuário (não exibida no JSON)
	IsAdmin   bool      `gorm:"default:false;not null" json:"is_admin"`     // Indica se o usuário é admin
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`           // Data de criação do registro
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`           // Data de atualização do registro
}

// Definição do repositório para 'Usuarios', com um campo DB para o acesso ao banco de dados.
type UsuariosRepository struct {
	DB *gorm.DB
}

// Função construtora que cria uma instância de 'UsuariosRepository' a partir de um banco GORM.
func NewUsuariosRepository(db *gorm.DB) *UsuariosRepository {
	log.Println("Repositório Usuários criado!")
	return &UsuariosRepository{DB: db}
}

// Função responsável por criar um novo usuário no banco de dados.
// Ela recebe um ponteiro para a struct 'Usuarios' e a insere na tabela correspondente.
func (repo *UsuariosRepository) CreateUsuario(usuario *Usuarios) error {
	// Cria o novo usuário no banco
	result := repo.DB.Create(&usuario)

	// Verifica se ocorreu algum erro ao criar o usuário
	if result.Error != nil {
		log.Printf("Erro ao criar usuário: %v, Dados do Usuário: %v", result.Error, usuario) // Log de erro com contexto
		return errors.New("erro ao criar usuário")
	}
	log.Printf("Usuário criado com sucesso! Nome: %s, ID: %d", usuario.Nome, usuario.ID) // Log de sucesso
	return nil
}

// Função que busca um usuário pelo seu ID.
// Recebe o ID como string e retorna um ponteiro para a struct 'Usuarios' ou um erro, caso não encontre.
func (repo *UsuariosRepository) FindUsuarioById(usuarioId string) (*Usuarios, error) {
	var usuario Usuarios
	// Busca o usuário pelo ID
	result := repo.DB.Where("id = ?", usuarioId).First(&usuario)

	// Se ocorrer erro ao buscar o usuário, loga o erro e retorna um erro.
	if result.Error != nil {
		log.Printf("Erro ao buscar usuário por ID: %v, ID: %s", result.Error, usuarioId) // Log de erro com o ID
		return nil, fmt.Errorf("erro ao achar usuário com ID %s: %w", usuarioId, result.Error)
	}

	log.Printf("Usuário encontrado: %s (ID: %d)", usuario.Nome, usuario.ID) // Log de sucesso
	return &usuario, nil
}

// Função que lista todos os usuários do banco de dados, ordenando pelo nome.
// Retorna uma slice de 'Usuarios' ou um erro caso ocorra algum problema.
func (repo *UsuariosRepository) ListAllUsuarios() ([]Usuarios, error) {
	var allUsuarios []Usuarios
	// Seleciona os usuários e ordena pelo nome, retornando apenas os campos necessários
	result := repo.DB.Order("nome ASC").Find(&allUsuarios)

	// Verifica se ocorreu algum erro ao buscar os usuários
	if result.Error != nil {
		log.Printf("Erro ao buscar usuários: %v", result.Error) // Log de erro
		return nil, fmt.Errorf("erro ao buscar usuários: %w", result.Error)
	} else if result.RowsAffected == 0 {
		log.Println("Nenhum usuário encontrado.") // Log quando não encontra usuários
		return nil, errors.New("nenhum usuário encontrado")
	}

	log.Printf("Foram encontrados %d usuários.", len(allUsuarios)) // Log de sucesso
	return allUsuarios, nil
}

// Função para deletar um usuário pelo ID.
// Recebe o ID como string e remove o registro correspondente no banco de dados.
func (repo *UsuariosRepository) DeleteUsuarioById(usuarioId string) error {
	var usuario Usuarios
	// Deleta o usuário com o ID fornecido
	result := repo.DB.Where("id = ?", usuarioId).Delete(&usuario)

	// Verifica se ocorreu algum erro ao deletar o usuário
	if result.Error != nil {
		log.Printf("Erro ao deletar usuário (ID: %s): %v", usuarioId, result.Error) // Log de erro com o ID
		return fmt.Errorf("erro ao deletar usuário com ID %s: %w", usuarioId, result.Error)
	}

	log.Printf("Usuário com ID %s deletado com sucesso!", usuarioId) // Log de sucesso
	return nil
}

// Função que busca um usuário pelo nome. Retorna o usuário encontrado ou um erro.
func (repo *UsuariosRepository) GetUsuarioLogin(nome string) (*Usuarios, error) {
	var usuario Usuarios
	// Busca o usuário pelo nome
	result := repo.DB.Where("nome = ?", nome).First(&usuario)

	// Verifica se houve erro ao buscar o usuário
	if result.Error != nil {
		log.Printf("Erro ao buscar usuário pelo nome: %v, Nome: %s", result.Error, nome) // Log de erro com o nome
		return nil, fmt.Errorf("erro ao achar usuário com nome %s: %w", nome, result.Error)
	}

	log.Printf("Usuário encontrado: %s (ID: %d)", usuario.Nome, usuario.ID) // Log de sucesso
	return &usuario, nil
}
