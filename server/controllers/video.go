package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/helpers/pgx"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllVideos(c *gin.Context) {
	videos, err := h.service.GetAllVideos(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(videos, "Video berhasil ditemukan"))
}

func (h *Controller) GetVideoByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id video wajib diisi"))
		return
	}

	video, err := h.service.GetVideoByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(video, "Video berhasil ditemukan"))
}

func (h *Controller) CreateVideo(c *gin.Context) {
	var reqBody struct {
		Title       string `json:"title" binding:"required"`
		Link        string `json:"link" binding:"required"`
		Description string `json:"description"`
		Author      string `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.InsertVideoParams{
		Title:       reqBody.Title,
		Link:        reqBody.Link,
		Description: pgx.NewTextFromString(reqBody.Description),
		Author:      reqBody.Author,
	}

	if err := h.service.CreateVideo(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Video berhasil dibuat"))
}

func (h *Controller) UpdateVideo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id video wajib diisi"))
		return
	}

	var reqBody struct {
		Title       string `json:"title" binding:"required"`
		Link        string `json:"link" binding:"required"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	params := postgresql.UpdateVideoParams{
		ID:          id,
		Title:       reqBody.Title,
		Link:        reqBody.Link,
		Description: pgx.NewTextFromString(reqBody.Description),
	}

	if err := h.service.UpdateVideo(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Video berhasil diperbarui"))
}

func (h *Controller) DeleteVideo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id video wajib diisi"))
		return
	}

	if err := h.service.DeleteVideo(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Video berhasil dihapus"))
}
