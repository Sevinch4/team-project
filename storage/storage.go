package storage

import "teamProject/api/models"

type IStorage interface {
	Close()
	Branch() IBranchStorage
}

type IBranchStorage interface {
	Create(models.CreateBranch) (string, error)
	GetByID(id string) (models.Branch, error)
	GetList(models.GetListRequest) (models.BranchResponse, error)
	Update(models.UpdateBranch) (string, error)
	Delete(id string) error
}
