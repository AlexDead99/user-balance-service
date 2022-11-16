package api

import (
	"net/http"

	db "github.com/AlexDead99/user-balance-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

type productsParams struct {
	ProductId int32 `json:"product_id" binding:"required"`
	Amount    int32 `json:"amount" binding:"required,gte=1"`
}

//assume that user can only buy products (service_id == 1), not sell them
type createTransferRequest struct {
	UserId      int32             `json:"user_id" binding:"required"`
	Products    []*productsParams `json:"products" binding:"required"`
	ServiceId   int32             `json:"service_id" binding:"required"`
	Description string            `json:"description" binding:"required"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var req createTransferRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
}

type fulfilTransferRequest struct {
	TransferId int32  `json:"transfer_id" binding:"required"`
	Status     string `json:"status" binding:"required,oneof=Success Failed"`
}

func (server *Server) fulfilTransfer(ctx *gin.Context) {
	var req fulfilTransferRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updateTransfer := db.UpdateTransferParams{
		TransferID: req.TransferId,
		Status:     req.Status,
	}

	if req.Status == "Success" {
		transfer, err := server.store.UpdateTransfer(ctx, updateTransfer)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusOK, transfer)
	}

	// Update Failed Transactions
}
