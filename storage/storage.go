package storage

import "teamProject/api/models"

type IStorage interface {
	Close()
	StaffTarif() IStaffTarifRepo
	Staff() IStaffRepo
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
