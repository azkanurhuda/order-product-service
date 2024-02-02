package costumer

import (
	"github.com/azkanurhuda/costumer-service/internal/entity"
	"gorm.io/gorm"
)

type Costumer struct {
	db *gorm.DB
}

func (p Costumer) GetLatestCostumer(db *gorm.DB) (*entity.Costumer, error) {
	var product entity.Costumer
	err := db.Order("id desc").First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p Costumer) Create(db *gorm.DB, product *entity.Costumer) error {
	return db.Create(product).Error
}

func (p Costumer) ReadByID(db *gorm.DB, id int) (*entity.Costumer, error) {
	var product entity.Costumer
	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p Costumer) ReadAll(db *gorm.DB) ([]entity.Costumer, error) {
	var products []entity.Costumer
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil

}

func (p Costumer) Update(db *gorm.DB, product *entity.Costumer) error {
	return db.Model(&entity.Costumer{}).Where("id = ?", product.ID).Updates(product).Error
}

func (p Costumer) Delete(db *gorm.DB, productID int) error {
	return db.Delete(&entity.Costumer{}, productID).Error
}

func NewCostumer(db *gorm.DB) *Costumer {
	return &Costumer{
		db: db,
	}
}
