package controllers

import (
	common "alter-io-go/helpers/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetRoles(c *gin.Context) {
	roles, err := h.service.GetRoles(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(roles, "Role berhasil ditemukan"))
}
