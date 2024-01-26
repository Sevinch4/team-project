package storage

import "teamProject/api/models"

type IStorage interface {
	Close()
	Branch() IBranchStorage
	Sale() ISaleStorage
	Transaction() ITransactionStorage
}

type IBranchStorage interface {
	Create(models.CreateBranch) (string, error)
	GetByID(id string) (models.Branch, error)
	GetList(models.GetListRequest) (models.BranchResponse, error)
	Update(models.UpdateBranch) (string, error)
	Delete(id string) error
}

type ISaleStorage interface {
	Create(models.CreateSale) (string, error)
	GetByID(id string) (models.Sale, error)
	GetList(models.GetListRequest) (models.SaleResponse, error)
	Update(sale models.UpdateSale) (string, error)
	Delete(id string) error
}

type ITransactionStorage interface {
	Create(models.CreateTransaction) (string, error)
	GetByID(id string) (models.Transaction, error)
	GetList(models.GetListRequest) (models.TransactionResponse, error)
	Update(models.UpdateTransaction) (string, error)
	Delete(id string) error
}
