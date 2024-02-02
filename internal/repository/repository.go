package repository

import (
	"github.com/azkanurhuda/costumer-service/internal/repository/costumer"
	"gorm.io/gorm"
)

type Repository struct {
	db       *gorm.DB
	Costumer Costumer
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:       db,
		Costumer: costumer.NewCostumer(db),
	}
}
