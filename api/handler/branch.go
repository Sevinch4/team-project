package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"teamProject/api/models"
)

func (h Handler) CreateBranch(c *gin.Context) {
	branch := models.CreateBranch{}

	if err := c.ShouldBindJSON(&branch); err != nil {
		handleResponse(c, "error is while reading body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storage.Branch().Create(branch)
	if err != nil {
		handleResponse(c, "error is while creating branch", http.StatusInternalServerError, err.Error())
		return
	}

	createdBranch, err := h.storage.Branch().GetByID(id)
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusCreated, createdBranch)
}

func (h Handler) GetBranch(c *gin.Context) {
	uid := c.Param("id")

	branch, err := h.storage.Branch().GetByID(uid)
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, branch)
}

func (h Handler) GetBranchList(c *gin.Context) {
	var (
		page, limit int
		search      string
		err         error
	)

	pageStr := c.DefaultQuery("page", "1")
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		handleResponse(c, "error is while converting page", http.StatusBadRequest, err.Error())
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		handleResponse(c, "error is while converting limit", http.StatusBadRequest, err.Error())
		return
	}

	search = c.Query("search")

	branches, err := h.storage.Branch().GetList(models.GetListRequest{
		Page:   page,
		Limit:  limit,
		Search: search,
	})

	handleResponse(c, "", http.StatusOK, branches)
}

func (h Handler) UpdateBranch(c *gin.Context) {
	uid := c.Param("id")

	branch := models.UpdateBranch{}
	if err := c.ShouldBindJSON(&branch); err != nil {
		handleResponse(c, "error is wile reading from body", http.StatusBadRequest, err.Error())
		return
	}

	branch.ID = uid
	id, err := h.storage.Branch().Update(branch)
	if err != nil {
		handleResponse(c, "error is while updating branch", http.StatusInternalServerError, err.Error())
		return
	}

	updatedBranch, err := h.storage.Branch().GetByID(id)
	if err != nil {
		handleResponse(c, "error is while getting by id", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, updatedBranch)
}

func (h Handler) DeleteBranch(c *gin.Context) {
	uid := c.Param("id")

	if err := h.storage.Branch().Delete(uid); err != nil {
		handleResponse(c, "error is while delteing branch", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "", http.StatusOK, "branch deleted!")
}
