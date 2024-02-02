package model

import "time"

type ProductRequest struct {
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
	Stock int     `json:"stock,omitempty"`
}

type ProductResponse struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Price     float64   `json:"price,omitempty"`
	Stock     int       `json:"stock,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
