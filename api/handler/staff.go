package handler

import (
	"net/http"
	"strconv"
	"teamProject/api/models"
	"teamProject/pkg/check"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CreateStaff(c *gin.Context) {
	staff := models.CreateStaff{}

	if err := c.ShouldBindJSON(&staff); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Staff().Create(staff)

	if err != nil {
		handleResponse(c, "error while creating staff ", http.StatusInternalServerError, err.Error())
		return
	}

	createdStaffTarif, err := h.storage.Staff().StaffByID(models.PrimaryKey{ID: id})
	if err != nil {
		handleResponse(c, "error while getting by ID", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, createdStaffTarif)
}

func (h *Handler) GetStaff(c *gin.Context) {
	uid := c.Param("id")

	staffTarif, err := h.storage.Staff().StaffByID(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting staff  by ID", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, staffTarif)
}

func (h *Handler) GetStaffList(c *gin.Context) {
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

	response, err := h.storage.Staff().GetStaffTList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})
	if err != nil {
		handleResponse(c, "error while getting staff list", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, response)
}

func (h *Handler) UpdateStaff(c *gin.Context) {
	uid := c.Param("id")

	staff := models.UpdateStaff{}
	if err := c.ShouldBindJSON(&staff); err != nil {
		handleResponse(c, "error while reading from body", http.StatusBadRequest, err.Error())
		return
	}

	staff.ID = uid
	if _, err := h.storage.Staff().UpdateStaff(staff); err != nil {
		handleResponse(c, "error while updating staff ", http.StatusInternalServerError, err.Error())
		return
	}

	updatedStaff, err := h.storage.Staff().StaffByID(models.PrimaryKey{ID: uid})
	if err != nil {
		handleResponse(c, "error while getting by ID", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, updatedStaff)
}

func (h *Handler) DeleteStaff(c *gin.Context) {
	uid := c.Param("id")

	if err := h.storage.Staff().DeleteStaff(uid); err != nil {
		handleResponse(c, "error while deleting staff ", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "staff tariff deleted")
}

func (h Handler) UpdateStaffPassword(c *gin.Context) {
	updateStaffPassword := models.UpdateStaffPassword{}

	if err := c.ShouldBindJSON(&updateStaffPassword); err != nil {
		handleResponse(c, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	uid, err := uuid.Parse(c.Param("id"))
	if err != nil {
		handleResponse(c, "error while parsing uuid", http.StatusBadRequest, err.Error())
		return
	}

	updateStaffPassword.ID = uid.String()

	oldPassword, err := h.storage.Staff().GetPassword(updateStaffPassword.ID)
	if err != nil {
		handleResponse(c, "error while getting password by id", http.StatusInternalServerError, err.Error())
		return
	}

	if oldPassword != updateStaffPassword.OldPassword {
		handleResponse(c, "old password is not correct", http.StatusBadRequest, "old password is not correct")
		return
	}

	if err = check.ValidatePassword(updateStaffPassword.NewPassword); err != nil {
		handleResponse(c, "new password is weak", http.StatusBadRequest, err.Error())
		return
	}

	if err = h.storage.Staff().UpdatePassword(updateStaffPassword); err != nil {
		handleResponse(c, "error while updating staff password by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, "password successfully updated")
}
