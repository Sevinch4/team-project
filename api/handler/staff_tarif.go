package handler

import (
	"net/http"
	"strconv"
	"teamProject/api/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateStaffTarif(c *gin.Context) {
	staffTarif := models.CreateStaffTarif{}

	if err := c.ShouldBindJSON(&staffTarif); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.StaffTarif().Create(staffTarif)
	if err != nil {
		handleResponse(c, "error while creating staff tariff", http.StatusInternalServerError, err.Error())
		return
	}

	createdStaffTarif, err := h.storage.StaffTarif().GetStaffTarifByID(models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error while getting by ID", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, createdStaffTarif)
}

func (h *Handler) GetStaffTarif(c *gin.Context) {
	uid := c.Param("id")

	staffTarif, err := h.storage.StaffTarif().GetStaffTarifByID(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting staff tariff by ID", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, staffTarif)
}

func (h *Handler) GetStaffTarifList(c *gin.Context) {
	var (
		page, limit int
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	search := c.Query("search")

	response, err := h.storage.StaffTarif().GetStaffTarifList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting staff tariff list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, response)
}

func (h *Handler) UpdateStaffTarif(c *gin.Context) {
	uid := c.Param("id")

	sTarif := models.UpdateStaffTarif{}
	if err := c.ShouldBindJSON(&sTarif); err != nil {
		handleResponse(c, "error while reading from body", http.StatusBadRequest, err.Error())
		return
	}

	sTarif.ID = uid
	if _, err := h.storage.StaffTarif().UpdateStaffTarif(sTarif); err != nil {
		handleResponse(c, "error while updating staff tariff", http.StatusInternalServerError, err.Error())
		return
	}

	updatedStaffTarif, err := h.storage.StaffTarif().GetStaffTarifByID(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting by ID", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, updatedStaffTarif)
}

func (h *Handler) DeleteStaffTarif(c *gin.Context) {
	uid := c.Param("id")

	if err := h.storage.StaffTarif().DeleteStaffTarif(uid); err != nil {
		handleResponse(c, "error while deleting staff tariff", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "staff tariff deleted")
}
