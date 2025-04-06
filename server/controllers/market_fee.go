package controllers

import (
	common "alter-io-go/helpers/http"
	"alter-io-go/helpers/pgx"
	"alter-io-go/repositories/postgresql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (h *Controller) GetAllMarketFees(c *gin.Context) {
	fees, err := h.service.GetAllMarketFees(c)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(fees, "Data retribusi pasar berhasil ditemukan"))
}

func (h *Controller) GetMarketFeeByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id retribusi wajib diisi"))
		return
	}

	fee, err := h.service.GetMarketFeeByID(c, id)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(fee, "Data retribusi pasar berhasil ditemukan"))
}

func (h *Controller) GetMarketFeesByMarket(c *gin.Context) {
	marketID := c.Param("marketId")
	if marketID == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id pasar wajib diisi"))
		return
	}

	fees, err := h.service.GetMarketFeesByMarket(c, marketID)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(fees, "Data retribusi pasar berhasil ditemukan"))
}

func (h *Controller) GetMarketFeesByYear(c *gin.Context) {
	yearStr := c.Param("year")
	if yearStr == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("tahun wajib diisi"))
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("tahun harus berupa angka"))
		return
	}

	fees, err := h.service.GetMarketFeesByYear(c, int32(year))
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(fees, "Data retribusi pasar berhasil ditemukan"))
}

func (h *Controller) GetMarketFeesBySemesterAndYear(c *gin.Context) {
	semester := c.Param("semester")
	if semester == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("semester wajib diisi"))
		return
	}

	yearStr := c.Param("year")
	if yearStr == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("tahun wajib diisi"))
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("tahun harus berupa angka"))
		return
	}

	params := postgresql.FindMarketFeesBySemesterAndYearParams{
		Semester: semester,
		Year:     int32(year),
	}

	fees, err := h.service.GetMarketFeesBySemesterAndYear(c, params)
	if err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(fees, "Data retribusi pasar berhasil ditemukan"))
}

func (h *Controller) CreateMarketFee(c *gin.Context) {
	var reqBody struct {
		MarketID                 string  `json:"market_id" binding:"required"`
		NumPermanentKiosks       int32   `json:"num_permanent_kiosks" binding:"required"`
		NumNonPermanentKiosks    int32   `json:"num_non_permanent_kiosks" binding:"required"`
		PermanentKioskRevenue    float64 `json:"permanent_kiosk_revenue" binding:"required"`
		NonPermanentKioskRevenue float64 `json:"non_permanent_kiosk_revenue" binding:"required"`
		CollectionStatus         string  `json:"collection_status" binding:"required"`
		Description              string  `json:"description"`
		Semester                 string  `json:"semester" binding:"required"`
		Year                     int32   `json:"year" binding:"required"`
		Author                   string  `json:"author" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	// Konversi nilai float ke Numeric menggunakan Scan
	var permRevenue pgtype.Numeric
	permRevenue.Valid = true
	_ = permRevenue.Scan(reqBody.PermanentKioskRevenue)

	var nonPermRevenue pgtype.Numeric
	nonPermRevenue.Valid = true
	_ = nonPermRevenue.Scan(reqBody.NonPermanentKioskRevenue)

	params := postgresql.InsertMarketFeeParams{
		MarketID:                 reqBody.MarketID,
		NumPermanentKiosks:       reqBody.NumPermanentKiosks,
		NumNonPermanentKiosks:    reqBody.NumNonPermanentKiosks,
		PermanentKioskRevenue:    permRevenue,
		NonPermanentKioskRevenue: nonPermRevenue,
		CollectionStatus:         reqBody.CollectionStatus,
		Description:              pgx.NewTextFromString(reqBody.Description),
		Semester:                 reqBody.Semester,
		Year:                     reqBody.Year,
		Author:                   reqBody.Author,
	}

	if err := h.service.CreateMarketFee(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusCreated, common.NewSuccessCreatedResponse("Data retribusi pasar berhasil dibuat"))
}

func (h *Controller) UpdateMarketFee(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id retribusi wajib diisi"))
		return
	}

	var reqBody struct {
		MarketID                 string  `json:"market_id" binding:"required"`
		NumPermanentKiosks       int32   `json:"num_permanent_kiosks" binding:"required"`
		NumNonPermanentKiosks    int32   `json:"num_non_permanent_kiosks" binding:"required"`
		PermanentKioskRevenue    float64 `json:"permanent_kiosk_revenue" binding:"required"`
		NonPermanentKioskRevenue float64 `json:"non_permanent_kiosk_revenue" binding:"required"`
		CollectionStatus         string  `json:"collection_status" binding:"required"`
		Description              string  `json:"description"`
		Semester                 string  `json:"semester" binding:"required"`
		Year                     int32   `json:"year" binding:"required"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse(err.Error()))
		return
	}

	// Konversi nilai float ke Numeric menggunakan Scan
	var permRevenue pgtype.Numeric
	permRevenue.Valid = true
	_ = permRevenue.Scan(reqBody.PermanentKioskRevenue)

	var nonPermRevenue pgtype.Numeric
	nonPermRevenue.Valid = true
	_ = nonPermRevenue.Scan(reqBody.NonPermanentKioskRevenue)

	params := postgresql.UpdateMarketFeeParams{
		ID:                       id,
		MarketID:                 reqBody.MarketID,
		NumPermanentKiosks:       reqBody.NumPermanentKiosks,
		NumNonPermanentKiosks:    reqBody.NumNonPermanentKiosks,
		PermanentKioskRevenue:    permRevenue,
		NonPermanentKioskRevenue: nonPermRevenue,
		CollectionStatus:         reqBody.CollectionStatus,
		Description:              pgx.NewTextFromString(reqBody.Description),
		Semester:                 reqBody.Semester,
		Year:                     reqBody.Year,
	}

	if err := h.service.UpdateMarketFee(c, params); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Data retribusi pasar berhasil diperbarui"))
}

func (h *Controller) DeleteMarketFee(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, common.NewBadRequestResponse("id retribusi wajib diisi"))
		return
	}

	if err := h.service.DeleteMarketFee(c, id); err != nil {
		resp := common.MapErrorToResponse(err)
		c.JSON(resp.Code, resp)
		return
	}

	c.JSON(http.StatusOK, common.NewSuccessDefaultResponse(nil, "Data retribusi pasar berhasil dihapus"))
}
