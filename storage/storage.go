package storage

import "teamProject/api/models"

type IStorage interface {
	Close()
	Category() ICategory
	Product() IProducts
	Branch() IBranchStorage
	Sale() ISaleStorage
	Transaction() ITransactionStorage
}

type ICategory interface {
	Create(models.CreateCategory) (string, error)
	GetByID(string) (models.Category, error)
	GetList(models.GetListRequest) (models.CategoryResponse, error)
	Update(models.UpdateCategory) (string, error)
	Delete(string) error
}

type IProducts interface {
	Create(models.CreateProduct) (string, error)
	GetByID(string) (models.Product, error)
	GetList(models.GetListRequest) (models.ProductResponse, error)
	Update(models.UpdateProduct) (string, error)
	Delete(string) error
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
