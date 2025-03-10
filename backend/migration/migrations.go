package migration

import (
	"log"

	"github.com/Turgho/barberfast/backend/models/repositories"
	"gorm.io/gorm"
)

func InitMigrations(db *gorm.DB) {
	err := db.AutoMigrate(
		&repositories.Usuarios{},
		&repositories.Servicos{},
		&repositories.Agendamentos{},
	)

	if err != nil {
		log.Fatal("erro ao fazer as migrações:", err)
		return
	}

	log.Println("Migrações de tabelas feitas!")
}
