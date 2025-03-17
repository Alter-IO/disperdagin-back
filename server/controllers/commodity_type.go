package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllCommodityTypes(c *gin.Context) {
	commodityTypes, err := h.service.GetAllCommodityTypes(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(commodityTypes, "Jenis komoditas berhasil ditemukan"))
}

func (h *Controller) GetCommodityTypeByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := errors.New("ID is required")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	commodityType, err := h.service.GetCommodityTypeByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(commodityType, "Jenis komoditas berhasil ditemukan"))
}

func (h *Controller) CreateCommodityType(c *gin.Context) {
	var reqBody struct {
		Description string `json:"description" binding:"required"`
		Author      string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	params := postgresql.InsertCommodityTypeParams{
		Description: reqBody.Description,
		Author:      reqBody.Author,
	}

	if err := h.service.CreateCommodityType(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Jenis komoditas berhasil dibuat"))
}

func (h *Controller) UpdateCommodityType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := errors.New("ID is required")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	var reqBody struct {
		Description string `json:"description" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		err := errors.New("Invalid request body")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	params := postgresql.UpdateCommodityTypeParams{
		ID:          id,
		Description: reqBody.Description,
	}

	if err := h.service.UpdateCommodityType(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Jenis komoditas berhasil diupdate"))
}

func (h *Controller) DeleteCommodityType(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		err := errors.New("ID is required")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	if err := h.service.DeleteCommodityType(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Jenis komoditas berhasil dihapus"))
}
