package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/helpers/logger"
	"alter-io-go/repositories/postgresql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllCommodities(c *gin.Context) {
	commodities, err := h.service.GetAllCommodities(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(commodities, "Komoditas berhasil ditemukan"))
}

func (h *Controller) GetLatestCommodities(c *gin.Context) {
	commodities, err := h.service.GetLatestCommodities(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(commodities, "Komoditas terbaru berhasil ditemukan"))
}

func (h *Controller) GetCommoditiesByType(c *gin.Context) {
	typeID := c.Param("typeId")
	if typeID == "" {
		err := errors.New("Tipe komoditas ID diperlukan")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	commodities, err := h.service.GetCommoditiesByType(c, typeID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(commodities, "Komoditas berhasil ditemukan"))
}

func (h *Controller) GetCommodityByID(c *gin.Context) {
	commodityID := c.Param("id")
	if commodityID == "" {
		err := errors.New("Komoditas ID diperlukan")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	commodity, err := h.service.GetCommodityByID(c, commodityID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(commodity, "Komoditas berhasil ditemukan"))
}

func (h *Controller) CreateCommodity(c *gin.Context) {
	reqBody := new(postgresql.InsertCommodityParams)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.MapErrorToResponse(err))
		return
	}

	logger.Get().With().Info("CreateCommodity", "reqBody", reqBody)

	if err := h.service.CreateCommodity(c, *reqBody); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Komoditas berhasil dibuat"))
}

func (h *Controller) UpdateCommodity(c *gin.Context) {
	commodityID := c.Param("id")
	if commodityID == "" {
		err := errors.New("Komoditas ID diperlukan")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	reqBody := new(postgresql.UpdateCommodityParams)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Set the ID from the URL parameter
	reqBody.ID = commodityID

	rowsAffected, err := h.service.UpdateCommodity(c, *reqBody)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Komoditas berhasil diperbarui"))
}

func (h *Controller) DeleteCommodity(c *gin.Context) {
	commodityID := c.Param("id")
	if commodityID == "" {
		err := errors.New("Komoditas ID diperlukan")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	rowsAffected, err := h.service.DeleteCommodity(c, commodityID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Komoditas berhasil dihapus"))
}
