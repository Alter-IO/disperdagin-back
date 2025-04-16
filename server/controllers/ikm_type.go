package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/helpers/pgx"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllIKMTypes(c *gin.Context) {
	ikmTypes, err := h.service.GetAllIKMTypes(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(ikmTypes, "Jenis IKM berhasil ditemukan"))
}

func (h *Controller) GetIKMTypeByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id jenis IKM wajib diisi"))
		return
	}

	ikmType, err := h.service.GetIKMTypeByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(ikmType, "Jenis IKM berhasil ditemukan"))
}

func (h *Controller) GetIKMTypesByInfoType(c *gin.Context) {
	infoType := c.Param("infoType")
	if infoType == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("tipe informasi wajib diisi"))
		return
	}

	ikmTypes, err := h.service.GetIKMTypesByInfoType(c, infoType)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(ikmTypes, "Jenis IKM berhasil ditemukan"))
}

func (h *Controller) CreateIKMType(c *gin.Context) {
	var reqBody struct {
		DocumentName   string `json:"document_name" binding:"required"`
		FileUrl        string `json:"file_url" binding:"required"`
		PublicInfoType string `json:"public_info_type" binding:"required"`
		Description    string `json:"description"`
		Author         string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertIKMTypeParams{
		DocumentName:   reqBody.DocumentName,
		FileUrl:        reqBody.FileUrl,
		PublicInfoType: reqBody.PublicInfoType,
		Description:    pgx.NewTextFromString(reqBody.Description),
		Author:         reqBody.Author,
	}

	if err := h.service.CreateIKMType(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Jenis IKM berhasil dibuat"))
}

func (h *Controller) UpdateIKMType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id jenis IKM wajib diisi"))
		return
	}

	var reqBody struct {
		DocumentName   string `json:"document_name" binding:"required"`
		FileUrl        string `json:"file_url" binding:"required"`
		PublicInfoType string `json:"public_info_type" binding:"required"`
		Description    string `json:"description"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdateIKMTypeParams{
		ID:             id,
		DocumentName:   reqBody.DocumentName,
		FileUrl:        reqBody.FileUrl,
		PublicInfoType: reqBody.PublicInfoType,
		Description:    pgx.NewTextFromString(reqBody.Description),
	}

	if err := h.service.UpdateIKMType(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Jenis IKM berhasil diperbarui"))
}

func (h *Controller) DeleteIKMType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id jenis IKM wajib diisi"))
		return
	}

	if err := h.service.DeleteIKMType(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Jenis IKM berhasil dihapus"))
}
