package api

import (
	"net/http"
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
// @Param body body <blank> true "reportBody"
// @Success      200  {object}  <blank>
// @Router       /users/report [post]
func (server *Server) CreateUserReport(ctx *gin.Context) {}

// ShowAccount godoc
// @Summary      Info about succeeded transfers for current month
// @Description  Info about succeeded transfers for current month
// @Tags         reports
// @Accept       json
// @Produce      json
// @Param body body MonthReportRequest true "transfer"
// @Success      200  {object}  <blank>
// @Router       /report [post]
type MonthReportRequest struct {
	Date time.Time `form:"date" binding:"required" time_format:"2022-01-02"`
}
type MonthReportResponse struct {
	Link string `json:"link"`
}

func (server *Server) CreateMonthReport(ctx *gin.Context) {
	var req MonthReportRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	params := db.GeneralReportParams{
		CreatedAt:   req.Date,
		CreatedAt_2: req.Date,
	}
	transfers, err := server.store.GeneralReport(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	reports := make(map[string]float32, 0)

	for _, transfer := range transfers {
		reports[transfer.name] += transfer.TotalPrice
	}

}
