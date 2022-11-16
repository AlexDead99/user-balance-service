package api

import (
	"net/http"
	"strconv"

	db "github.com/AlexDead99/user-balance-service/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner   string  `json:"owner" binding:"required"`
	Balance float32 `json:"balance" binding:"gte=1"`
}

// ShowAccount godoc
// @Summary      Create an account
// @Description  Create user's account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param body body createAccountRequest true "user"
// @Success      200  {object}  db.Accounts
// @Router       /accounts [post]
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:   req.Owner,
		Balance: req.Balance,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type updateAccountRequest struct {
	Amount float32 `json:"amount" binding:"required"`
}

// ShowAccount godoc
// @Summary      Update account's balance
// @Description  Update user's account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Param body body updateAccountRequest true "user"
// @Success      200  {object}  db.UpdateUserBalanceTxResult
// @Router       /accounts/{id} [put]
func (server *Server) updateAccount(ctx *gin.Context) {
	var req updateAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id := ctx.Param("id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	updateParams := db.UpdateUserBalanceTxParams{
		UserId: int32(userId),
		Amount: req.Amount,
	}

	result, err := server.store.UpdateUserBalanceTx(ctx, updateParams)

	ctx.JSON(http.StatusOK, result)
}

// ShowAccount godoc
// @Summary      Get user's account
// @Description  Get user's account
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  db.Accounts
// @Router       /accounts/{id} [get]
func (server *Server) getAccountInfo(ctx *gin.Context) {

	id := ctx.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	existedUser, err := server.store.GetAccount(ctx, int32(userId))
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, existedUser)
}
