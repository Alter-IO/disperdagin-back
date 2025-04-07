package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllLegalDocTypes(c *gin.Context) {
	docTypes, err := h.service.GetAllLegalDocTypes(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(docTypes, "Tipe dokumen hukum berhasil ditemukan"))
}

func (h *Controller) GetLegalDocTypeByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id tipe dokumen wajib diisi"))
		return
	}

	docType, err := h.service.GetLegalDocTypeByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(docType, "Tipe dokumen hukum berhasil ditemukan"))
}

func (h *Controller) CreateLegalDocType(c *gin.Context) {
	var reqBody struct {
		Description string `json:"description" binding:"required"`
		Author      string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertLegalDocTypeParams{
		Description: reqBody.Description,
		Author:      reqBody.Author,
	}

	if err := h.service.CreateLegalDocType(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Tipe dokumen hukum berhasil dibuat"))
}

func (h *Controller) UpdateLegalDocType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id tipe dokumen wajib diisi"))
		return
	}

	var reqBody struct {
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdateLegalDocTypeParams{
		ID:          id,
		Description: reqBody.Description,
	}

	if err := h.service.UpdateLegalDocType(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Tipe dokumen hukum berhasil diperbarui"))
}

func (h *Controller) DeleteLegalDocType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id tipe dokumen wajib diisi"))
		return
	}

	if err := h.service.DeleteLegalDocType(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Tipe dokumen hukum berhasil dihapus"))
}
