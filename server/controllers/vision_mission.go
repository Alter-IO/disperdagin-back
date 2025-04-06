package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllVisionMissions(c *gin.Context) {
	visionMissions, err := h.service.GetAllVisionMissions(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(visionMissions, "Visi dan misi berhasil ditemukan"))
}

func (h *Controller) GetLatestVisionMission(c *gin.Context) {
	visionMission, err := h.service.GetLatestVisionMission(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(visionMission, "Visi dan misi terbaru berhasil ditemukan"))
}

func (h *Controller) GetVisionMissionByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id visi dan misi wajib diisi"))
		return
	}

	visionMission, err := h.service.GetVisionMissionByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(visionMission, "Visi dan misi berhasil ditemukan"))
}

func (h *Controller) CreateVisionMission(c *gin.Context) {
	var reqBody postgresql.InsertVisionMissionParams
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	if err := h.service.CreateVisionMission(c, reqBody); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Visi dan misi berhasil dibuat"))
}

func (h *Controller) UpdateVisionMission(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id visi dan misi wajib diisi"))
		return
	}

	var reqBody postgresql.UpdateVisionMissionParams
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	if err := h.service.UpdateVisionMission(c, id, reqBody); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Visi dan misi berhasil diperbarui"))
}

func (h *Controller) DeleteVisionMission(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id visi dan misi wajib diisi"))
		return
	}

	if err := h.service.DeleteVisionMission(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Visi dan misi berhasil dihapus"))
}
