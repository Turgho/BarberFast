package repositories

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Definição do modelo de dados para a tabela 'usuarios'
// 'Usuarios' mapeia as colunas do banco de dados com as propriedades do struct.
type Usuarios struct {
	ID        string    `gorm:"primaryKey;size:36" json:"id"` // Armazenando UUID como string
	Nome      string    `gorm:"size:255;not null" json:"nome"`
	Email     string    `gorm:"uniqueIndex;size:255;not null" json:"email"`
	Telefone  string    `gorm:"size:14;not null" json:"telefone"`
	Senha     string    `gorm:"not null" json:"senha"`
	IsAdmin   bool      `gorm:"default:false;not null" json:"is_admin"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// GerarUUID para gerar o UUID antes de salvar
func (u *Usuarios) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String() // Gera um UUID antes de criar o usuário
	return nil
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
	// Verifica se o email já existe no banco antes de tentar criar o usuário
	var existingUsuario Usuarios

	result := repo.DB.Where("email = ?", usuario.Email).First(&existingUsuario)

	if result.Error == nil {
		// Se o resultado for nil, significa que o email já está no banco
		log.Printf("erro: email já cadastrado. email: %s", usuario.Email)
		return fmt.Errorf("o email %s já está cadastrado", usuario.Email)

	} else if result.Error != gorm.ErrRecordNotFound {
		// Caso tenha ocorrido outro erro
		log.Printf("Erro ao verificar o email: %v", result.Error)
		return fmt.Errorf("erro ao verificar o email: %v", result.Error)
	}

	// Cria o novo usuário no banco
	result = repo.DB.Create(&usuario)

	// Verifica se ocorreu algum erro ao criar o usuário
	if result.Error != nil {
		log.Printf("Erro ao criar usuário: %v, Dados do Usuário: %v", result.Error, usuario) // Log de erro com contexto
		return fmt.Errorf("erro ao criar usuário: %v", result.Error)
	}

	log.Printf("Usuário criado com sucesso! Nome: %s, ID: %s", usuario.Nome, usuario.ID) // Log de sucesso
	return nil
}

// Função que busca um usuário pelo seu ID.
// Recebe o ID como string e retorna um ponteiro para a struct 'Usuarios' ou um erro, caso não encontre.
func (repo *UsuariosRepository) FindUsuarioById(usuarioId string) (*Usuarios, error) {
	var usuario Usuarios

	// Busca o usuário pelo ID
	result := repo.DB.
		Select("nome", "email", "tefelone").
		Where("id = ?", usuarioId).
		First(&usuario)

	// Se ocorrer erro ao buscar o usuário, loga o erro e retorna um erro.
	if result.Error != nil {
		log.Printf("Erro ao buscar usuário por ID: %v, ID: %s", result.Error, usuarioId) // Log de erro com o ID
		return nil, fmt.Errorf("erro ao achar usuário com ID %s: %w", usuarioId, result.Error)
	}

	log.Printf("Usuário encontrado: %s (ID: %s)", usuario.Nome, usuario.ID) // Log de sucesso
	return &usuario, nil
}

// Função que lista todos os usuários do banco de dados, ordenando pelo nome.
// Retorna uma slice de 'Usuarios' ou um erro caso ocorra algum problema.
func (repo *UsuariosRepository) ListAllUsuarios() ([]Usuarios, error) {
	var allUsuarios []Usuarios

	// Busca os usuários no banco e retorna apenas os campos necessários
	result := repo.DB.
		Model(&Usuarios{}).
		Select("nome", "email", "telefone").
		Order("nome ASC").
		Find(&allUsuarios)

	// Se houver erro, retorna
	if result.Error != nil {
		log.Printf("Erro ao buscar usuários: %v", result.Error)
		return nil, fmt.Errorf("erro ao buscar usuários: %w", result.Error)
	}

	log.Printf("Foram encontrados %d usuários.", len(allUsuarios))
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

	log.Printf("Usuário com ID %s deletaddo com sucesso!", usuarioId) // Log de sucesso
	return nil
}

// Função que busca um usuário pelo nome. Retorna o usuário encontrado ou um erro.
func (repo *UsuariosRepository) GetUsuarioLogin(email string) (*Usuarios, error) {
	var usuario Usuarios
	// Busca o usuário pelo nome
	result := repo.DB.Where("email = ?", email).First(&usuario)

	// Verifica se houve erro ao buscar o usuário
	if result.Error != nil {
		log.Printf("Erro ao buscar usuário pelo email: %v, Email: %s", result.Error, email) // Log de erro com o nome
		return nil, fmt.Errorf("erro ao achar usuário com email %s: %w", email, result.Error)
	}

	log.Printf("Usuário encontrado: %s (ID: %s)", usuario.Nome, usuario.ID) // Log de sucesso
	return &usuario, nil
}
