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

	r.POST("/category", h.CreateCategory)
	r.GET("/category/:id", h.GetCategory)
	r.GET("/categories", h.GetCategoryList)
	r.PUT("/category/:id", h.UpdateCategory)
	r.DELETE("/category/:id", h.DeleteCategory)

	r.POST("/branch", h.CreateBranch)
	r.GET("/branch/:id", h.GetBranch)
	r.GET("/branches", h.GetBranchList)
	r.PUT("/branch/:id", h.UpdateBranch)
	r.DELETE("/branch/:id", h.DeleteBranch)

	r.POST("/sale", h.CreateSale)
	r.GET("/sale/:id", h.GetSale)
	r.GET("/sales", h.GetSaleList)
	r.PUT("/sale/:id", h.UpdateSale)
	r.DELETE("/sale/:id", h.DeleteSale)

	r.POST("/transaction", h.CreateTransaction)
	r.GET("/transaction/:id", h.GetTransaction)
	r.GET("/transactions", h.GetTransactionList)
	r.PUT("/transaction/:id", h.UpdateTransaction)
	r.DELETE("/transaction/:id", h.DeleteTransaction)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
	return r

}
