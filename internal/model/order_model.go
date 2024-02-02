package model

import "time"

type OrderRequest struct {
	CustomerID int     `json:"customer_id,omitempty"`
	ProductID  int     `json:"product_id,omitempty"`
	Quantity   int     `json:"quantity,omitempty"`
	Total      float64 `json:"total,omitempty"`
}

type OrderResponse struct {
	ID         int       `json:"id,omitempty"`
	CustomerID int       `json:"customer_id,omitempty"`
	ProductID  int       `json:"product_id,omitempty"`
	Quantity   int       `json:"quantity,omitempty"`
	Total      float64   `json:"total,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt  time.Time `json:"updated_at,omitempty"`
}

type OrderFullResponse struct {
	ID        int              `json:"id,omitempty"`
	Customer  CostumerResponse `json:"customer,omitempty"`
	Product   ProductResponse  `json:"product,omitempty"`
	Quantity  int              `json:"quantity,omitempty"`
	Total     float64          `json:"total,omitempty"`
	CreatedAt time.Time        `json:"created_at,omitempty"`
	UpdatedAt time.Time        `json:"updated_at,omitempty"`
}

type CostumerResponse struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ProductResponse struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Stock     int       `json:"stock,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
