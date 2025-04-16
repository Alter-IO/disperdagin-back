package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/helpers/pgx"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllPublicInfo(c *gin.Context) {
	info, err := h.service.GetAllPublicInfo(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(info, "Informasi publik berhasil ditemukan"))
}

func (h *Controller) GetPublicInfoByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id informasi publik wajib diisi"))
		return
	}

	info, err := h.service.GetPublicInfoByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(info, "Informasi publik berhasil ditemukan"))
}

func (h *Controller) GetPublicInfoByType(c *gin.Context) {
	infoType := c.Param("infoType")
	if infoType == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("tipe informasi publik wajib diisi"))
		return
	}

	info, err := h.service.GetPublicInfoByType(c, infoType)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(info, "Informasi publik berhasil ditemukan"))
}

func (h *Controller) CreatePublicInfo(c *gin.Context) {
	var reqBody struct {
		DocumentName   string `json:"document_name" binding:"required"`
		FileUrl        string `json:"file_name" binding:"required"`
		PublicInfoType string `json:"public_info_type" binding:"required"`
		Description    string `json:"description"`
		Author         string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertPublicInfoParams{
		DocumentName:   reqBody.DocumentName,
		FileUrl:        reqBody.FileUrl,
		PublicInfoType: reqBody.PublicInfoType,
		Description:    pgx.NewTextFromString(reqBody.Description),
		Author:         reqBody.Author,
	}

	if err := h.service.CreatePublicInfo(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Informasi publik berhasil dibuat"))
}

func (h *Controller) UpdatePublicInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id informasi publik wajib diisi"))
		return
	}

	var reqBody struct {
		DocumentName   string `json:"document_name" binding:"required"`
		FileUrl        string `json:"file_name" binding:"required"`
		PublicInfoType string `json:"public_info_type" binding:"required"`
		Description    string `json:"description"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdatePublicInfoParams{
		ID:             id,
		DocumentName:   reqBody.DocumentName,
		FileUrl:        reqBody.FileUrl,
		PublicInfoType: reqBody.PublicInfoType,
		Description:    pgx.NewTextFromString(reqBody.Description),
	}

	if err := h.service.UpdatePublicInfo(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Informasi publik berhasil diperbarui"))
}

func (h *Controller) DeletePublicInfo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id informasi publik wajib diisi"))
		return
	}

	if err := h.service.DeletePublicInfo(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Informasi publik berhasil dihapus"))
}
