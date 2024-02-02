package repository

import (
	"github.com/azkanurhuda/product-service/internal/repository/product"
	"gorm.io/gorm"
)

type Repository struct {
	db      *gorm.DB
	Product Product
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:      db,
		Product: product.NewProduct(db),
	}
}
