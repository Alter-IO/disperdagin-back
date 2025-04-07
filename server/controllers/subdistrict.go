package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllSubdistricts(c *gin.Context) {
	subdistricts, err := h.service.GetAllSubdistricts(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(subdistricts, "Kecamatan berhasil ditemukan"))
}

func (h *Controller) GetSubdistrictByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id kecamatan wajib diisi"))
		return
	}

	subdistrict, err := h.service.GetSubdistrictByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(subdistrict, "Kecamatan berhasil ditemukan"))
}

func (h *Controller) CreateSubdistrict(c *gin.Context) {
	var reqBody struct {
		Name   string `json:"name" binding:"required"`
		Author string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertSubdistrictParams{
		Name:   reqBody.Name,
		Author: reqBody.Author,
	}

	if err := h.service.CreateSubdistrict(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Kecamatan berhasil dibuat"))
}

func (h *Controller) UpdateSubdistrict(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id kecamatan wajib diisi"))
		return
	}

	var reqBody struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdateSubdistrictParams{
		ID:   id,
		Name: reqBody.Name,
	}

	if err := h.service.UpdateSubdistrict(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Kecamatan berhasil diperbarui"))
}

func (h *Controller) DeleteSubdistrict(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id kecamatan wajib diisi"))
		return
	}

	if err := h.service.DeleteSubdistrict(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Kecamatan berhasil dihapus"))
}
