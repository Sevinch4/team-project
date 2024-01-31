package storage

import (
	"context"
	"teamProject/api/models"
)

type IStorage interface {
	Close()
	Category() ICategory
	Product() IProducts
	Branch() IBranchStorage
	Sale() ISaleStorage
	Transaction() ITransactionStorage
}

type ICategory interface {
	Create(context.Context, models.CreateCategory) (string, error)
	GetByID(context.Context, string) (models.Category, error)
	GetList(context.Context, models.GetListRequest) (models.CategoryResponse, error)
	Update(context.Context, models.UpdateCategory) (string, error)
	Delete(context.Context, string) error
}

type IProducts interface {
	Create(context.Context, models.CreateProduct) (string, error)
	GetByID(context.Context, string) (models.Product, error)
	GetList(context.Context, models.ProductGetListRequest) (models.ProductResponse, error)
	Update(context.Context, models.UpdateProduct) (string, error)
	Delete(context.Context, string) error
}

type IBranchStorage interface {
	Create(context.Context, models.CreateBranch) (string, error)
	GetByID(context.Context, string) (models.Branch, error)
	GetList(context.Context, models.GetListRequest) (models.BranchResponse, error)
	Update(context.Context, models.UpdateBranch) (string, error)
	Delete(context.Context, string) error
}

type ISaleStorage interface {
	Create(context.Context, models.CreateSale) (string, error)
	GetByID(context.Context, string) (models.Sale, error)
	GetList(context.Context, models.GetListRequest) (models.SaleResponse, error)
	Update(context.Context, models.UpdateSale) (string, error)
	Delete(context.Context, string) error
}

type ITransactionStorage interface {
	Create(context.Context, models.CreateTransaction) (string, error)
	GetByID(context.Context, string) (models.Transaction, error)
	GetList(context.Context, models.TransactionGetListRequest) (models.TransactionResponse, error)
	Update(context.Context, models.UpdateTransaction) (string, error)
	Delete(context.Context, string) error
}
