package repository

import (
	"github.com/azkanurhuda/order-product-service/internal/repository/order"
	"gorm.io/gorm"
)

type Repository struct {
	db    *gorm.DB
	Order Order
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db:    db,
		Order: order.NewOrder(db),
	}
}
