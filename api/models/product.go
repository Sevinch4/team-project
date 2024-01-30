package models

import "time"

type Product struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Barcode    int       `json:"barcode"`
	CategoryID string    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"-"`
}

type CreateProduct struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Barcode    int    `json:"barcode"`
	CategoryID string `json:"category_id"`
}

type UpdateProduct struct {
	ID         string    `json:"-"`
	Name       string    `json:"name"`
	Price      int       `json:"price"`
	Barcode    int       `json:"barcode"`
	CategoryID string    `json:"category_id"`
	UpdatedAt  time.Time `json:"-"`
}

type ProductResponse struct {
	Products []Product
	Count    int
}
