package controllers

import (
	"alter-io-go/domain"
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
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

func (h *Controller) GetDailyCommodities(c *gin.Context) {
	commodities, err := h.service.GetDailyCommodities(c)
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
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("tipe komoditas id diperlukan"))
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
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("tipe komoditas id diperlukan"))
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

func (h *Controller) CreateDailyCommodity(c *gin.Context) {
	reqBody := new(domain.CommodityDaily)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	if err := h.service.CreateDailyCommodity(c, *reqBody); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Komoditas harian berhasil dibuat"))
}

func (h *Controller) CreateCommodity(c *gin.Context) {
	reqBody := new(postgresql.InsertCommodityParams)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

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
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("komoditas id diperlukan"))
		return
	}

	reqBody := new(postgresql.UpdateCommodityParams)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	// Set the ID from the URL parameter
	reqBody.ID = commodityID

	if err := h.service.UpdateCommodity(c, *reqBody); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Komoditas berhasil diperbarui"))
}

func (h *Controller) DeleteCommodity(c *gin.Context) {
	commodityID := c.Param("id")
	if commodityID == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("komoditas id diperlukan"))
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
