package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAccounts(c *gin.Context) {
	users, err := h.service.GetUsers(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(users, "Akun berhasil ditemukan"))
}

func (h *Controller) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("user id is required"))
		return
	}

	user, err := h.service.GetUserByID(c, userID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(user, "Akun berhasil ditemukan"))
}

func (h *Controller) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("username is required"))
		return
	}

	user, err := h.service.GetUserByUsername(c, username)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(user, "Akun berhasil ditemukan"))
}

func (h *Controller) CreateUser(c *gin.Context) {
	reqBody := new(postgresql.InsertUserParams)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	if err := h.service.CreateUser(c, *reqBody); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Akun berhasil dibuat"))
}

func (h *Controller) UpdatePassword(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("username is required"))
		return
	}

	type UpdatePasswordRequest struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	reqBody := new(UpdatePasswordRequest)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	if err := h.service.UpdatePassword(c.Request.Context(), username, reqBody.OldPassword, reqBody.NewPassword); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Kata sandi berhasil diperbarui"))
}

func (h *Controller) ResetPassword(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("username is required"))
		return
	}

	type ResetPasswordReq struct {
		NewPassword string `json:"new_password" binding:"required"`
	}

	reqBody := new(ResetPasswordReq)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	err := h.service.ResetPassword(c.Request.Context(), userId, reqBody.NewPassword)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Kata sandi berhasil direset"))
}

func (h *Controller) DeleteUser(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("username is required"))
		return
	}

	if err := h.service.DeleteUser(c, username); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Akun berhasil dihapus"))
}
