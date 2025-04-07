package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllGreetings(c *gin.Context) {
	greetings, err := h.service.GetAllGreetings(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(greetings, "Sambutan berhasil ditemukan"))
}

func (h *Controller) GetLatestGreeting(c *gin.Context) {
	greeting, err := h.service.GetLatestGreeting(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(greeting, "Sambutan terbaru berhasil ditemukan"))
}

func (h *Controller) GetGreetingByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id sambutan wajib diisi"))
		return
	}

	greeting, err := h.service.GetGreetingByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(greeting, "Sambutan berhasil ditemukan"))
}

func (h *Controller) CreateGreeting(c *gin.Context) {
	var reqBody struct {
		Message string `json:"message" binding:"required"`
		Author  string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertGreetingParams{
		Message: reqBody.Message,
		Author:  reqBody.Author,
	}

	if err := h.service.CreateGreeting(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Sambutan berhasil dibuat"))
}

func (h *Controller) UpdateGreeting(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id sambutan wajib diisi"))
		return
	}

	var reqBody struct {
		Message string `json:"message" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdateGreetingParams{
		ID:      id,
		Message: reqBody.Message,
	}

	if err := h.service.UpdateGreeting(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Sambutan berhasil diperbarui"))
}

func (h *Controller) DeleteGreeting(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id sambutan wajib diisi"))
		return
	}

	if err := h.service.DeleteGreeting(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Sambutan berhasil dihapus"))
}
