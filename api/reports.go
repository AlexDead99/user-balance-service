package api

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	db "github.com/AlexDead99/user-balance-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

type MonthReportRequest struct {
	Date string `json:"date" binding:"required"`
}
type MonthReportResponse struct {
	Link   string             `json:"link"`
	Report map[string]float32 `json:"report"`
}

// ShowAccount godoc
// @Summary      Info about succeeded transfers for current month
// @Description  Info about succeeded transfers for current month
// @Tags         reports
// @Accept       json
// @Produce      json
// @Param body body MonthReportRequest true "transfer"
// @Success      200  {object} MonthReportResponse
// @Router       /report [post]
func (server *Server) CreateMonthReport(ctx *gin.Context) {
	var req MonthReportRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Can't convert to date")
		return
	}

	month := int32(date.Month())
	year := int32(date.Year())
	params := db.GeneralReportParams{
		CreatedAt:   month,
		CreatedAt_2: year,
	}

	transfers, err := server.store.GeneralReport(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	reports := make(map[string]float32, 0)

	for _, transfer := range transfers {
		reports[transfer.Name] += transfer.TotalPrice
	}

	fileName := strconv.Itoa(int(month)) + "-" + strconv.Itoa(date.Year()) + ".csv"
	pwd, err := os.Getwd()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fullPath := pwd + "/reports/" + fileName
	f, err := os.Create(fullPath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for key, value := range reports {
		err := w.Write([]string{fmt.Sprintf("%v", key), fmt.Sprintf("%v", value)})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
	}

	downloadPath := "http://localhost:3000/static/" + fileName

	ctx.JSON(http.StatusOK, MonthReportResponse{Link: downloadPath, Report: reports})
}
