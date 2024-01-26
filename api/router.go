package api

import (
	"teamProject/api/handler"
	"teamProject/storage"

	"github.com/gin-gonic/gin"
)

func New(storage storage.IStorage) *gin.Engine {
	h := handler.New(storage)

	r := gin.New()

	r.POST("/stafftarif", h.CreateStaffTarif)
	r.GET("/stafftarif/:id", h.GetStaffTarif)
	r.GET("/stafftarifs", h.GetStaffTarifList)
	r.PUT("/stafftarif/:id", h.UpdateStaffTarif)
	r.DELETE("/stafftarif/:id", h.DeleteStaffTarif)

	// Staff

	r.POST("/staff", h.CreateStaff)
	r.GET("/staff/:id", h.GetStaff)
	r.GET("/staffs", h.GetStaffList)
	r.PUT("/staff/:id", h.UpdateStaff)
	r.DELETE("/staff/:id", h.DeleteStaff)

	return r
}
