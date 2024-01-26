package postgres

import (
	"database/sql"
	"teamProject/api/models"
	"teamProject/storage"
)

type saleRepo struct {
	db *sql.DB
}

func NewSaleRepo(db *sql.DB) storage.ISaleStorage {
	return saleRepo{db: db}
}

func (s saleRepo) Create(sale models.CreateSale) (string, error) {
	return "", nil
}

func (s saleRepo) GetByID(id string) (models.Sale, error) {
	return models.Sale{}, nil
}

func (s saleRepo) GetList(request models.GetListRequest) (models.SaleResponse, error) {
	return models.SaleResponse{}, nil
}

func (s saleRepo) Update(sale models.UpdateSale) (string, error) {
	return "", nil
}

func (s saleRepo) Delete(id string) error {
	return nil
}
