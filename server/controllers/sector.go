package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetSectors(c *gin.Context) {
	sectors, err := h.service.GetSectors(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(sectors, "Sektor berhasil ditemukan"))
}

func (h *Controller) GetSector(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id sektor wajib diisi"))
		return
	}

	sector, err := h.service.GetSector(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(sector, "Sektor berhasil ditemukan"))
}

func (h *Controller) CreateSector(c *gin.Context) {
	var reqBody postgresql.InsertSectorParams
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	if err := h.service.CreateSector(c, reqBody); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Sektor berhasil dibuat"))
}

func (h *Controller) UpdateSector(c *gin.Context) {
	var reqBody postgresql.UpdateSectorParams
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id sektor wajib diisi"))
		return
	}

	if err := h.service.UpdateSector(c, id, reqBody); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Sektor berhasil diperbarui"))
}

func (h *Controller) DeleteSector(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id sektor wajib diisi"))
		return
	}

	if err := h.service.DeleteSector(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Sektor berhasil dihapus"))
}
