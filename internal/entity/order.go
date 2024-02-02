package entity

import "time"

type Order struct {
	ID         int       `gorm:"column:id;primaryKey"`
	CustomerID int       `gorm:"column:customer_id"`
	ProductID  int       `gorm:"column:product_id"`
	Quantity   int       `gorm:"column:quantity"`
	Total      float64   `gorm:"column:total"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}

func (u *Order) TableName() string {
	return "orders"
}
