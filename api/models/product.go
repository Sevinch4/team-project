package models

type Product struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Barcode    int    `json:"barcode"`
	CategoryID string `json:"category_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	DeletedAt  string `json:"-"`
}

type CreateProduct struct {
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Barcode    int    `json:"barcode"`
	CategoryID string `json:"category_id"`
}

type UpdateProduct struct {
	ID         string `json:"-"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Barcode    int    `json:"barcode"`
	CategoryID string `json:"category_id"`
	UpdatedAt  string `json:"-"`
}

type ProductResponse struct {
	Products []Product
	Count    int
}
