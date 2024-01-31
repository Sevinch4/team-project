package models

import (
	"time"
)

type Transaction struct {
	ID              string    `json:"id"`
	SaleID          string    `json:"sale_id"`
	StaffID         string    `json:"staff_id"`
	TransactionType string    `json:"transaction_type"`
	SourceType      string    `json:"source_type"`
	Amount          float32   `json:"amount"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       string    `json:"-"`
}

type CreateTransaction struct {
	SaleID          string  `json:"sale_id"`
	StaffID         string  `json:"staff_id"`
	TransactionType string  `json:"transaction_type"`
	SourceType      string  `json:"source_type"`
	Amount          float32 `json:"amount"`
	Description     string  `json:"description"`
}

type UpdateTransaction struct {
	ID              string    `json:"-"`
	SaleID          string    `json:"sale_id"`
	StaffID         string    `json:"staff_id"`
	TransactionType string    `json:"transaction_type"`
	SourceType      string    `json:"source_type"`
	Amount          float32   `json:"amount"`
	Description     string    `json:"description"`
	UpdatedAt       time.Time `json:"-"`
}

type TransactionResponse struct {
	Transactions []Transaction
	Count        int
}
