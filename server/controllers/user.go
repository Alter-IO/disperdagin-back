package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"errors"
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

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(users))
}

func (h *Controller) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		err := errors.New("User ID is required")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	user, err := h.service.GetUserByID(c, userID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(user))
}

func (h *Controller) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		err := errors.New("Username is required")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	user, err := h.service.GetUserByUsername(c, username)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(user))
}

func (h *Controller) CreateUser(c *gin.Context) {
	reqBody := new(postgresql.InsertUserParams)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if err := h.service.CreateUser(c, *reqBody); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse())
}

func (h *Controller) UpdatePassword(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		err := errors.New("Username is required")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	type UpdatePasswordRequest struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required"`
	}

	reqBody := new(UpdatePasswordRequest)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		err := errors.New("Invalid request body")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	if err := h.service.UpdatePassword(c.Request.Context(), username, reqBody.OldPassword, reqBody.NewPassword); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse("Password updated successfully"))
}

func (h *Controller) ResetPassword(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		err := errors.New("Username is required")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	// No request body needed as we'll generate a random password
	newPassword, err := h.service.ResetPassword(c.Request.Context(), username)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	// Return the generated password to the admin
	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(gin.H{
		"message":      "Password reset successfully",
		"new_password": newPassword,
	}))
}

func (h *Controller) DeleteUser(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		err := errors.New("Username is required")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	if err := h.service.DeleteUser(c, username); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse("User deleted successfully"))
}
