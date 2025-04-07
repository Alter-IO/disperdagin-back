package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllMarkets(c *gin.Context) {
	markets, err := h.service.GetAllMarkets(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(markets, "Pasar berhasil ditemukan"))
}

func (h *Controller) GetMarketByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id pasar wajib diisi"))
		return
	}

	market, err := h.service.GetMarketByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(market, "Pasar berhasil ditemukan"))
}

func (h *Controller) CreateMarket(c *gin.Context) {
	var reqBody struct {
		Name   string `json:"name" binding:"required"`
		Author string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertMarketParams{
		Name:   reqBody.Name,
		Author: reqBody.Author,
	}

	if err := h.service.CreateMarket(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Pasar berhasil dibuat"))
}

func (h *Controller) UpdateMarket(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id pasar wajib diisi"))
		return
	}

	var reqBody struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdateMarketParams{
		ID:   id,
		Name: reqBody.Name,
	}

	if err := h.service.UpdateMarket(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Pasar berhasil diperbarui"))
}

func (h *Controller) DeleteMarket(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id pasar wajib diisi"))
		return
	}

	if err := h.service.DeleteMarket(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Pasar berhasil dihapus"))
}
