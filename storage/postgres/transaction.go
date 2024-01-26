package postgres

import (
	"database/sql"
	"teamProject/api/models"
	"teamProject/storage"
)

type transactionRepo struct {
	db *sql.DB
}

func NewTransactionRepo(db *sql.DB) storage.ITransactionStorage {
	return transactionRepo{db: db}
}

func (t transactionRepo) Create(transaction models.CreateTransaction) (string, error) {
	return "", nil
}

func (t transactionRepo) GetByID(id string) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (t transactionRepo) GetList(request models.GetListRequest) (models.TransactionResponse, error) {
	return models.TransactionResponse{}, nil
}

func (t transactionRepo) Update(transaction models.UpdateTransaction) (string, error) {
	return "", nil
}

func (t transactionRepo) Delete(id string) error {
	return nil
}
