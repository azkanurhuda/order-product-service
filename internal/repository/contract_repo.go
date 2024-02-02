package repository

import (
	"github.com/azkanurhuda/order-product-service/internal/entity"
	"gorm.io/gorm"
)

type Order interface {
	Create(db *gorm.DB, product *entity.Order) error
	ReadByID(db *gorm.DB, id int) (*entity.Order, error)
	ReadAll(db *gorm.DB) ([]entity.Order, error)
	Update(db *gorm.DB, product *entity.Order) error
	Delete(db *gorm.DB, productID int) error
	GetLatestOrder(db *gorm.DB) (*entity.Order, error)
}
