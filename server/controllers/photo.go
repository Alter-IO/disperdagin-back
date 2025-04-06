package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/helpers/pgx"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllPhotos(c *gin.Context) {
	photos, err := h.service.GetAllPhotos(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(photos, "Foto berhasil ditemukan"))
}

func (h *Controller) GetPhotoByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id foto wajib diisi"))
		return
	}

	photo, err := h.service.GetPhotoByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(photo, "Foto berhasil ditemukan"))
}

func (h *Controller) GetPhotosByCategory(c *gin.Context) {
	categoryID := c.Param("categoryId")
	if categoryID == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id kategori wajib diisi"))
		return
	}

	photos, err := h.service.GetPhotosByCategory(c, categoryID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(photos, "Foto berhasil ditemukan"))
}

func (h *Controller) CreatePhoto(c *gin.Context) {
	var reqBody struct {
		CategoryID  string `json:"category_id" binding:"required"`
		Title       string `json:"title" binding:"required"`
		File        string `json:"file" binding:"required"`
		Description string `json:"description"`
		Author      string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertPhotoParams{
		CategoryID:  reqBody.CategoryID,
		Title:       reqBody.Title,
		File:        pgx.NewTextFromString(reqBody.File),
		Description: pgx.NewTextFromString(reqBody.Description),
		Author:      reqBody.Author,
	}

	if err := h.service.CreatePhoto(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Foto berhasil dibuat"))
}

func (h *Controller) UpdatePhoto(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id foto wajib diisi"))
		return
	}

	var reqBody struct {
		CategoryID  string `json:"category_id" binding:"required"`
		Title       string `json:"title" binding:"required"`
		File        string `json:"file"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdatePhotoParams{
		ID:          id,
		CategoryID:  reqBody.CategoryID,
		Title:       reqBody.Title,
		File:        pgx.NewTextFromString(reqBody.File),
		Description: pgx.NewTextFromString(reqBody.Description),
	}

	if err := h.service.UpdatePhoto(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Foto berhasil diperbarui"))
}

func (h *Controller) DeletePhoto(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id foto wajib diisi"))
		return
	}

	if err := h.service.DeletePhoto(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Foto berhasil dihapus"))
}
