package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllPublicInfoTypes(c *gin.Context) {
	infoTypes, err := h.service.GetAllPublicInfoTypes(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(infoTypes, "Jenis informasi publik berhasil ditemukan"))
}

func (h *Controller) GetPublicInfoTypeByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id jenis informasi publik wajib diisi"))
		return
	}

	infoType, err := h.service.GetPublicInfoTypeByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(infoType, "Jenis informasi publik berhasil ditemukan"))
}

func (h *Controller) CreatePublicInfoType(c *gin.Context) {
	var reqBody struct {
		Description string `json:"description" binding:"required"`
		Author      string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertPublicInfoTypeParams{
		Description: reqBody.Description,
		Author:      reqBody.Author,
	}

	if err := h.service.CreatePublicInfoType(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Jenis informasi publik berhasil dibuat"))
}

func (h *Controller) UpdatePublicInfoType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id jenis informasi publik wajib diisi"))
		return
	}

	var reqBody struct {
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdatePublicInfoTypeParams{
		ID:          id,
		Description: reqBody.Description,
	}

	if err := h.service.UpdatePublicInfoType(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Jenis informasi publik berhasil diperbarui"))
}

func (h *Controller) DeletePublicInfoType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id jenis informasi publik wajib diisi"))
		return
	}

	if err := h.service.DeletePublicInfoType(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Jenis informasi publik berhasil dihapus"))
}
