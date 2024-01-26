package storage

import "teamProject/api/models"

type IStorage interface {
	Close()
	StaffTarif() IStaffTarifRepo
}

type IStaffTarifRepo interface {
	Create(tarif models.CreateStaffTarif) (string, error)
	GetStaffTarifByID(models.PrimaryKey) (models.StaffTarif, error)
	GetStaffTarifList(req models.GetListRequest) (models.StaffTarifResponse, error)
	UpdateStaffTarif(models.UpdateStaffTarif) (string, error)
	DeleteStaffTarif(id string) error
}
