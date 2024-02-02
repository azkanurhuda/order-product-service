package repository

import (
	"github.com/azkanurhuda/product-service/internal/entity"
	"gorm.io/gorm"
)

type Product interface {
	Create(db *gorm.DB, product *entity.Product) error
	ReadByID(db *gorm.DB, id int) (*entity.Product, error)
	ReadAll(db *gorm.DB) ([]entity.Product, error)
	Update(db *gorm.DB, product *entity.Product) error
	Delete(db *gorm.DB, productID int) error
	GetLatestProduct(db *gorm.DB) (*entity.Product, error)
}
