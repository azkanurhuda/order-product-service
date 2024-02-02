package order

import (
	"github.com/azkanurhuda/order-product-service/internal/entity"
	"gorm.io/gorm"
)

type Order struct {
	db *gorm.DB
}

func (p Order) GetLatestOrder(db *gorm.DB) (*entity.Order, error) {
	var product entity.Order
	err := db.Order("id desc").First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p Order) Create(db *gorm.DB, product *entity.Order) error {
	return db.Create(product).Error
}

func (p Order) ReadByID(db *gorm.DB, id int) (*entity.Order, error) {
	var product entity.Order
	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p Order) ReadAll(db *gorm.DB) ([]entity.Order, error) {
	var products []entity.Order
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil

}

func (p Order) Update(db *gorm.DB, product *entity.Order) error {
	return db.Model(&entity.Order{}).Where("id = ?", product.ID).Updates(product).Error
}

func (p Order) Delete(db *gorm.DB, productID int) error {
	return db.Delete(&entity.Order{}, productID).Error
}

func NewOrder(db *gorm.DB) *Order {
	return &Order{
		db: db,
	}
}
