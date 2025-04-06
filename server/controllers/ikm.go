package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllIKMs(c *gin.Context) {
	ikms, err := h.service.GetAllIKMs(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(ikms, "Data IKM berhasil ditemukan"))
}

func (h *Controller) GetIKMByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id IKM wajib diisi"))
		return
	}

	ikm, err := h.service.GetIKMByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(ikm, "Data IKM berhasil ditemukan"))
}

func (h *Controller) GetIKMsByVillage(c *gin.Context) {
	villageID := c.Param("villageId")
	if villageID == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id desa wajib diisi"))
		return
	}

	ikms, err := h.service.GetIKMsByVillage(c, villageID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(ikms, "Data IKM berhasil ditemukan"))
}

func (h *Controller) GetIKMsByBusinessType(c *gin.Context) {
	businessType := c.Param("businessType")
	if businessType == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("tipe bisnis wajib diisi"))
		return
	}

	ikms, err := h.service.GetIKMsByBusinessType(c, businessType)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(ikms, "Data IKM berhasil ditemukan"))
}

func (h *Controller) CreateIKM(c *gin.Context) {
	var reqBody struct {
		Description  string `json:"description" binding:"required"`
		VillageID    string `json:"village_id" binding:"required"`
		BusinessType string `json:"business_type" binding:"required"`
		Author       string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertIKMParams{
		Description:  reqBody.Description,
		VillageID:    reqBody.VillageID,
		BusinessType: reqBody.BusinessType,
		Author:       reqBody.Author,
	}

	if err := h.service.CreateIKM(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Data IKM berhasil dibuat"))
}

func (h *Controller) UpdateIKM(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id IKM wajib diisi"))
		return
	}

	var reqBody struct {
		Description  string `json:"description" binding:"required"`
		VillageID    string `json:"village_id" binding:"required"`
		BusinessType string `json:"business_type" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdateIKMParams{
		ID:           id,
		Description:  reqBody.Description,
		VillageID:    reqBody.VillageID,
		BusinessType: reqBody.BusinessType,
	}

	if err := h.service.UpdateIKM(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Data IKM berhasil diperbarui"))
}

func (h *Controller) DeleteIKM(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id IKM wajib diisi"))
		return
	}

	if err := h.service.DeleteIKM(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Data IKM berhasil dihapus"))
}
