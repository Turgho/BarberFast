package repositories

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Usuarios struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Nome      string    `gorm:"size:255;not null" json:"nome"`
	Email     string    `gorm:"uniqueIndex;size:255;not null" json:"email"`
	Telefone  string    `gorm:"size:14;not null" json:"telefone"`
	Senha     string    `gorm:"not null" json:"senha"`                  // Omitido no JSON, mas não no banco
	IsAdmin   bool      `gorm:"default:false;not null" json:"is_admin"` // Default false, visível no JSON
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UsuariosRepository struct {
	DB *gorm.DB
}

func NewUsuariosRepository(db *gorm.DB) *UsuariosRepository {
	return &UsuariosRepository{DB: db}
}

func (repo *UsuariosRepository) CreateUsuario(usuario *Usuarios) error {
	// Cria um novo cliente
	result := repo.DB.Create(&usuario)

	if result.Error != nil {
		return fmt.Errorf("erro ao criar usuario: %v", result.Error)
	}
	return nil
}

func (repo *UsuariosRepository) FindUsuarioById(usuarioId string) (*Usuarios, error) {
	var usuario Usuarios
	result := repo.DB.Where("id = ?", usuarioId).First(&usuario)

	if result.Error != nil {
		return nil, fmt.Errorf("erro ao achar usuario: %v", result.Error)
	}

	return &usuario, nil
}

func (repo *UsuariosRepository) ListAllUsuarios() ([]Usuarios, error) {
	var allUsuarios []Usuarios
	result := repo.DB.Order("nome ASC").Find(&allUsuarios)

	if result.Error != nil {
		return nil, fmt.Errorf("erro ao achar os clientes: %v", result.Error)
	} else if result.RowsAffected == 0 {
		return nil, fmt.Errorf("nenhum usuário encontrado: %v", nil)
	}

	return allUsuarios, nil
}

func (repo *UsuariosRepository) DeleteUsuarioById(usuarioId string) error {
	var usuario Usuarios
	result := repo.DB.Where("id = ?", usuarioId).Delete(usuario)

	if result.Error != nil {
		return fmt.Errorf("erro ao deletar usuário: %v", result.Error)
	}
	return nil
}

func (repo *UsuariosRepository) GetUsuarioLogin(nome string) (*Usuarios, error) {
	var usuario Usuarios
	result := repo.DB.
		Where("nome = ?", nome).
		First(&usuario)

	if result.Error != nil {
		return nil, errors.New("erro ao achar usuario: " + result.Error.Error())
	}
	return &usuario, nil
}
