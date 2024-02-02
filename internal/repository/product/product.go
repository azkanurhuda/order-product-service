package product

import (
	"github.com/azkanurhuda/product-service/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	db *gorm.DB
}

func (p Product) GetLatestProduct(db *gorm.DB) (*entity.Product, error) {
	var product entity.Product
	err := db.Order("id desc").First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p Product) Create(db *gorm.DB, product *entity.Product) error {
	return db.Create(product).Error
}

func (p Product) ReadByID(db *gorm.DB, id int) (*entity.Product, error) {
	var product entity.Product
	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p Product) ReadAll(db *gorm.DB) ([]entity.Product, error) {
	var products []entity.Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil

}

func (p Product) Update(db *gorm.DB, product *entity.Product) error {
	return db.Model(&entity.Product{}).Where("id = ?", product.ID).Updates(product).Error
}

func (p Product) Delete(db *gorm.DB, productID int) error {
	return db.Delete(&entity.Product{}, productID).Error
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		db: db,
	}
}
