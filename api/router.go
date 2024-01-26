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

	return r

}
