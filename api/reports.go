package api

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"

	db "github.com/AlexDead99/user-balance-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
// @Summary      Info about succeeded user's transfers
// @Description  Info about succeeded user's transfers
// @Tags         reports
// @Accept       json
// @Produce      json
// @Param body body MonthReportRequest true "reportBody"
// @Success      200  {object} MonthReportResponse
// @Router       /users/report [post]
func (server *Server) CreateUserReport(ctx *gin.Context) {}

type MonthReportRequest struct {
	Date string `json:"date" binding:"required"`
}
type MonthReportResponse struct {
	Link string `json:"link"`
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
	fmt.Println(date, err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Can't convert to date")
		return
	}

	params := db.GeneralReportParams{
		CreatedAt:   date,
		CreatedAt_2: date,
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

	fileName := date.String() + ".csv"
	f, err := os.Create("../" + fileName)
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

	ctx.JSON(http.StatusOK, MonthReportResponse{Link: fileName})
}
