package controllers

import (
	"alter-io-go/domain"
	common "alter-io-go/helpers/http"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) LoginUser(c *gin.Context) {
	reqBody := new(domain.LoginReq)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	auth, err := h.service.Login(c, *reqBody)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(auth, "Anda Berhasil Masuk!"))
}
