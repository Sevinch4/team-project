package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"teamProject/api/models"
)

func (h Handler) CreateSale(c *gin.Context) {
	sale := models.CreateSale{}
	if err := c.ShouldBindJSON(&sale); err != nil {
		handleResponse(c, "error is while reading from body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Sale().Create(sale)
	if err != nil {
		handleResponse(c, "error is while creating sale", http.StatusInternalServerError, err.Error())
		return
	}

	createdBranch, err := h.storage.Sale().GetByID(id)
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, createdBranch)
}

func (h Handler) GetSale(c *gin.Context) {

}

func (h Handler) GetSaleList(c *gin.Context) {

}

func (h Handler) UpdateSale(c *gin.Context) {

}

func (h Handler) DeleteSale(c *gin.Context) {

}
