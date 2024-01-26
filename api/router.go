package api

import (
	"github.com/gin-gonic/gin"
	"teamProject/api/handler"
	"teamProject/storage"
)

func New(storage storage.IStorage) *gin.Engine {
	h := handler.New(storage)

	r := gin.New()

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

	return r

}
