package api

import (
	"net/http"

	db "github.com/AlexDead99/user-balance-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

//assume that user can only buy products (service_id == 1), not sell them
type createTransferRequest struct {
	UserId      int32                `json:"user_id" binding:"required"`
	Products    []*db.ProductsParams `json:"products" binding:"required"`
	ServiceId   int32                `json:"service_id" binding:"required"`
	Description string               `json:"description" binding:"required"`
}

// ShowAccount godoc
// @Summary      Create transfer
// @Description  Create transfer
// @Tags         transfers
// @Accept       json
// @Produce      json
// @Param body body createTransferRequest true "transfer"
// @Success      200  {object}  db.TransferTxResult
// @Router       /transfers [post]
func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	amountCheck := false
	for _, product := range req.Products {
		if product.Amount < 1 {
			amountCheck = true
			break
		}
	}
	if amountCheck == true {
		ctx.JSON(http.StatusBadRequest, "Invalid product's amount")
		return
	}

	createParams := db.TransferTxParams{
		UserId:      req.UserId,
		ServiceId:   1,
		Description: req.Description,
		Products:    req.Products,
	}
	result, err := server.store.TransferTx(ctx, createParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, result)
}

type fulfilTransferRequest struct {
	TransferId int32  `json:"transfer_id" binding:"required"`
	Status     string `json:"status" binding:"required,oneof=Success Failed"`
}

type fulfilTransferResponse struct {
	Status bool `json:"status"`
}

// ShowAccount godoc
// @Summary      Fulfil transfer
// @Description  Fulfil transfer
// @Tags         transfers
// @Accept       json
// @Produce      json
// @Param body body fulfilTransferRequest true "transfer"
// @Success      200  {object}  fulfilTransferResponse
// @Router       /transfers [put]
func (server *Server) fulfilTransfer(ctx *gin.Context) {
	var req fulfilTransferRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	transfer, err := server.store.GetTransfer(ctx, req.TransferId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if transfer.Status != "Pending" {
		ctx.JSON(http.StatusBadRequest, "You can't update status for non-pending transfer")
		return
	}

	if req.Status == "Success" {
		updateTransfer := db.UpdateTransferParams{
			TransferID: req.TransferId,
			Status:     req.Status,
		}
		_, err := server.store.UpdateTransfer(ctx, updateTransfer)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, fulfilTransferResponse{Status: true})
	}

	txParams := db.DeleteTransferTxParams{
		TransferId: req.TransferId,
	}
	result, err := server.store.DeleteTransferTx(ctx, txParams)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, fulfilTransferResponse{Status: result.Success})
}
