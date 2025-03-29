package controllers

import (
	"alter-io-go/domain"
	"alter-io-go/helpers/derrors"
	common "alter-io-go/helpers/http"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (h *Controller) UploadFile(c *gin.Context) {
	var reqBody domain.UploadReq
	if err := c.ShouldBind(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	uploadInfo, err := h.service.UploadFile(c, reqBody)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	if err := c.SaveUploadedFile(reqBody.File, filepath.Join(".", uploadInfo.FileURL)); err != nil {
		errType := derrors.WrapErrorf(err, derrors.ErrorCodeUnknown, "%s", err.Error())
		resp := common.MapErrorToResponse(errType)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewCreatedResponseWithData(uploadInfo))
}
