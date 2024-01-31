package storage

import (
	"context"
	"teamProject/api/models"
)

type IStorage interface {
	Close()
	StaffTarif() IStaffTarifRepo
	Staff() IStaffRepo
	Repository() IRepositoryRepo
	Category() ICategory
	Product() IProducts
	Branch() IBranchStorage
	Sale() ISaleStorage
	Transaction() ITransactionStorage
}

type IStaffTarifRepo interface {
	Create(tarif models.CreateStaffTarif) (string, error)
	GetStaffTarifByID(models.PrimaryKey) (models.StaffTarif, error)
	GetStaffTarifList(req models.GetListRequest) (models.StaffTarifResponse, error)
	UpdateStaffTarif(models.UpdateStaffTarif) (string, error)
	DeleteStaffTarif(id string) error
}

type IStaffRepo interface {
	Create(models.CreateStaff) (string, error)
	StaffByID(models.PrimaryKey) (models.Staff, error)
	GetStaffTList(models.GetListRequest) (models.StaffsResponse, error)
	UpdateStaff(models.UpdateStaff) (string, error)
	DeleteStaff(id string) error
	GetPassword(id string) (string, error)
	UpdatePassword(models.UpdateStaffPassword) error
}

type IRepositoryRepo interface {
	Create(models.CreateRepository) (string, error)
	GetByID(models.PrimaryKey) (models.Repository, error)
	GetList(models.GetListRequest) (models.RepositoriesResponse, error)
	Update(models.UpdateRepository) (string, error)
	Delete(string) error
}

type IBasketRepo interface {
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
