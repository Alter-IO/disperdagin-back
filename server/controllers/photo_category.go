package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllPhotoCategories(c *gin.Context) {
	categories, err := h.service.GetAllPhotoCategories(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(categories, "Kategori foto berhasil ditemukan"))
}

func (h *Controller) GetPhotoCategoryByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id kategori foto wajib diisi"))
		return
	}

	category, err := h.service.GetPhotoCategoryByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(category, "Kategori foto berhasil ditemukan"))
}

func (h *Controller) CreatePhotoCategory(c *gin.Context) {
	var reqBody struct {
		Category string `json:"category" binding:"required"`
		Author   string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertPhotoCategoryParams{
		Category: reqBody.Category,
		Author:   reqBody.Author,
	}

	if err := h.service.CreatePhotoCategory(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Kategori foto berhasil dibuat"))
}

func (h *Controller) UpdatePhotoCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id kategori foto wajib diisi"))
		return
	}

	var reqBody struct {
		Category string `json:"category" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdatePhotoCategoryParams{
		ID:       id,
		Category: reqBody.Category,
	}

	if err := h.service.UpdatePhotoCategory(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Kategori foto berhasil diperbarui"))
}

func (h *Controller) DeletePhotoCategory(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id kategori foto wajib diisi"))
		return
	}

	if err := h.service.DeletePhotoCategory(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Kategori foto berhasil dihapus"))
}
