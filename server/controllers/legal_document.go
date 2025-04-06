package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/helpers/pgx"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllLegalDocuments(c *gin.Context) {
	documents, err := h.service.GetAllLegalDocuments(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(documents, "Dokumen hukum berhasil ditemukan"))
}

func (h *Controller) GetLegalDocumentByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id dokumen hukum wajib diisi"))
		return
	}

	document, err := h.service.GetLegalDocumentByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(document, "Dokumen hukum berhasil ditemukan"))
}

func (h *Controller) GetLegalDocumentsByType(c *gin.Context) {
	docType := c.Param("docType")
	if docType == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("tipe dokumen wajib diisi"))
		return
	}

	documents, err := h.service.GetLegalDocumentsByType(c, docType)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(documents, "Dokumen hukum berhasil ditemukan"))
}

func (h *Controller) CreateLegalDocument(c *gin.Context) {
	var reqBody struct {
		DocumentName string `json:"document_name" binding:"required"`
		FileName     string `json:"file_name" binding:"required"`
		DocumentType string `json:"document_type" binding:"required"`
		Description  string `json:"description"`
		Author       string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertLegalDocumentParams{
		DocumentName: reqBody.DocumentName,
		FileName:     reqBody.FileName,
		DocumentType: reqBody.DocumentType,
		Description:  pgx.NewTextFromString(reqBody.Description),
		Author:       reqBody.Author,
	}

	if err := h.service.CreateLegalDocument(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Dokumen hukum berhasil dibuat"))
}

func (h *Controller) UpdateLegalDocument(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id dokumen hukum wajib diisi"))
		return
	}

	var reqBody struct {
		DocumentName string `json:"document_name" binding:"required"`
		FileName     string `json:"file_name" binding:"required"`
		DocumentType string `json:"document_type" binding:"required"`
		Description  string `json:"description"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdateLegalDocumentParams{
		ID:           id,
		DocumentName: reqBody.DocumentName,
		FileName:     reqBody.FileName,
		DocumentType: reqBody.DocumentType,
		Description:  pgx.NewTextFromString(reqBody.Description),
	}

	if err := h.service.UpdateLegalDocument(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Dokumen hukum berhasil diperbarui"))
}

func (h *Controller) DeleteLegalDocument(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id dokumen hukum wajib diisi"))
		return
	}

	if err := h.service.DeleteLegalDocument(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Dokumen hukum berhasil dihapus"))
}
