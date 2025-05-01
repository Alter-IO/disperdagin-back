package controllers

import (
	"alter-io-go/helpers/derrors"
	common "alter-io-go/helpers/http"
	"alter-io-go/repositories/postgresql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetAllEmployees(c *gin.Context) {
	employees, err := h.service.GetAllEmployees(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(employees, "Karyawan berhasil ditemukan"))
}

func (h *Controller) GetActiveEmployees(c *gin.Context) {
	employees, err := h.service.GetActiveEmployees(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(employees, "Karyawan aktif berhasil ditemukan"))
}

func (h *Controller) GetEmployeesByPosition(c *gin.Context) {
	position := c.Param("position")
	if position == "" {
		err := errors.New("Jabatan karyawan diperlukan")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	employees, err := h.service.GetEmployeesByPosition(c, position)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(employees, "Karyawan berhasil ditemukan"))
}

func (h *Controller) GetEmployeeByID(c *gin.Context) {
	employeeID := c.Param("id")
	if employeeID == "" {
		err := errors.New("ID karyawan diperlukan")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	employee, err := h.service.GetEmployeeByID(c, employeeID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(employee, "Karyawan berhasil ditemukan"))
}

func (h *Controller) CreateEmployee(c *gin.Context) {
	reqBody := new(postgresql.InsertEmployeeParams)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	if err := h.service.CreateEmployee(c, *reqBody); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Karyawan berhasil dibuat"))
}

func (h *Controller) UpdateEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	if employeeID == "" {
		err := errors.New("ID karyawan diperlukan")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	reqBody := new(postgresql.UpdateEmployeeParams)
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	// Set the ID from the URL parameter
	reqBody.ID = employeeID

	rowsAffected, err := h.service.UpdateEmployee(c, *reqBody)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	if rowsAffected == 0 {
		notFoundErr := derrors.NewErrorf(derrors.ErrorCodeNotFound, "karyawan tidak ditemukan")
		resp := common.MapErrorToResponse(notFoundErr)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Karyawan berhasil diperbarui"))
}

func (h *Controller) DeleteEmployee(c *gin.Context) {
	employeeID := c.Param("id")
	if employeeID == "" {
		err := errors.New("ID karyawan diperlukan")
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	rowsAffected, err := h.service.DeleteEmployee(c, employeeID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	if rowsAffected == 0 {
		notFoundErr := derrors.NewErrorf(derrors.ErrorCodeNotFound, "karyawan tidak ditemukan")
		resp := common.MapErrorToResponse(notFoundErr)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Karyawan berhasil dihapus"))
}

func (h *Controller) GetEmployeePositions(c *gin.Context) {
	positions, err := h.service.GetEmployeePositions(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(positions, "Jabatan berhasil ditemukan"))
}
