package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllVillages(c *gin.Context) {
	villages, err := h.service.GetAllVillages(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(villages, "Desa berhasil ditemukan"))
}

func (h *Controller) GetVillageByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id desa wajib diisi"))
		return
	}

	village, err := h.service.GetVillageByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(village, "Desa berhasil ditemukan"))
}

func (h *Controller) GetVillagesBySubdistrict(c *gin.Context) {
	subdistrictID := c.Param("subdistrictId")
	if subdistrictID == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id kecamatan wajib diisi"))
		return
	}

	villages, err := h.service.GetVillagesBySubdistrict(c, subdistrictID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(villages, "Desa berhasil ditemukan"))
}

func (h *Controller) CreateVillage(c *gin.Context) {
	var reqBody struct {
		SubdistrictID string `json:"subdistrict_id" binding:"required"`
		Name          string `json:"name" binding:"required"`
		Author        string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertVillageParams{
		SubdistrictID: reqBody.SubdistrictID,
		Name:          reqBody.Name,
		Author:        reqBody.Author,
	}

	if err := h.service.CreateVillage(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Desa berhasil dibuat"))
}

func (h *Controller) UpdateVillage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id desa wajib diisi"))
		return
	}

	var reqBody struct {
		SubdistrictID string `json:"subdistrict_id" binding:"required"`
		Name          string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdateVillageParams{
		ID:            id,
		SubdistrictID: reqBody.SubdistrictID,
		Name:          reqBody.Name,
	}

	if err := h.service.UpdateVillage(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Desa berhasil diperbarui"))
}

func (h *Controller) DeleteVillage(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id desa wajib diisi"))
		return
	}

	if err := h.service.DeleteVillage(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Desa berhasil dihapus"))
}
