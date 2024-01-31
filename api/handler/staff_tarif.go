package handler

import (
	"net/http"
	"strconv"
	"teamProject/api/models"

	"github.com/gin-gonic/gin"
)

// CreateStaffTarif godoc
// @Router       /stafftarif [POST]
// @Summary      Create a new stafftarif
// @Description  create a new stafftarif
// @Tags         stafftarif
// @Accept       json
// @Produce      json
// @Param 		 staffTarif body models.CreateStaffTarif false "staffTarif"
// @Success      200  {object}  models.StaffTarif
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
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

// GetStaffTarif godoc
// @Router       /stafftarif/{id} [GET]
// @Summary      Get stafftarif by id
// @Description  get stafftarif by id
// @Tags         stafftarif
// @Accept       json
// @Produce      json
// @Param 		 id path string true "stafftarif_id"
// @Success      200  {object}  models.StaffTarif
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) GetStaffTarif(c *gin.Context) {
	uid := c.Param("id")

	staffTarif, err := h.storage.StaffTarif().GetStaffTarifByID(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting staff tariff by ID", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, staffTarif)
}

// GetStaffTarifList godoc
// @Router       /stafftarifs [GET]
// @Summary      Get stafftarif list
// @Description  get stafftarif list
// @Tags         stafftarif
// @Accept       json
// @Produce      json
// @Param 		 page query string false "page"
// @Param 		 limit query string false "limit"
// @Param 		 search query string false "search"
// @Success      200  {object}  models.StaffTarifResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
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

// UpdateStaffTarif godoc
// @Router       /stafftarif/{id} [PUT]
// @Summary      Update stafftarif
// @Description  get stafftarif
// @Tags         stafftarif
// @Accept       json
// @Produce      json
// @Param 		 id path string true "stafftarif_id"
// @Param 		 stafftarif body models.UpdateStaffTarif false "stafftarif"
// @Success      200  {object}  models.StaffTarif
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
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

// DeleteStaffTarif godoc
// @Router       /stafftarif/{id} [DELETE]
// @Summary      Delete stafftarif
// @Description  delete stafftarif
// @Tags         stafftarif
// @Accept       json
// @Produce      json
// @Param 		 id path string true "stafftarif_id"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) DeleteStaffTarif(c *gin.Context) {
	uid := c.Param("id")

	if err := h.storage.StaffTarif().DeleteStaffTarif(uid); err != nil {
		handleResponse(c, "error while deleting staff tariff", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "staff tariff deleted")
}
