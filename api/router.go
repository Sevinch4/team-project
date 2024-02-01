package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "teamProject/api/docs"
	"teamProject/api/handler"
	"teamProject/storage"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
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

	// Repository 
	
	r.POST("/repository", h.CreateRepository)
	r.GET("/repository/:id", h.GetRepository)
	r.GET("/repositories", h.GetRepositoryList)
	r.PUT("/repository/:id", h.UpdateRepository)
	r.DELETE("/repository/:id", h.DeleteRepository)

	// Basket

	r.POST("/basket", h.CreateBasket)
	r.GET("/basket/:id", h.GetBasket)
	r.GET("/baskets", h.GetBasket)
	r.PUT("/basket/:id", h.UpdateBasket)
	r.DELETE("/basket/:id", h.DeleteBasket)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
	return r
}
