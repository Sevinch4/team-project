package api

import (
	"github.com/gin-gonic/gin"
	"teamProject/api/handler"
	"teamProject/storage"
)

func New(storage storage.IStorage) *gin.Engine {
	h := handler.New(storage)

	r := gin.New()

	r.POST("/stafftarif", h.CreateStaffTarif)
	r.GET("/stafftarif/:id", h.GetStaffTarif)
	r.GET("/stafftarifs", h.GetStaffTarifList)
	r.PUT("/stafftarif/:id", h.UpdateStaffTarif) // Add a leading slash here
	r.DELETE("/stafftarif/:id", h.DeleteStaffTarif)

	return r
}
