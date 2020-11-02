package api

import (
	"net/http"
	db "pismo/db/sqlc"

	"github.com/gin-gonic/gin"
)

const (
	DEBIT     = 1
	CREDIT    = 2
	WITHDRAWL = 3
	PAYMENT   = 4
)

type createTransactionRequest struct {
	AccountID     int32   `json:"account_id" binding:"required"`
	OperationType int16   `json:"operation_type_id" binding:"required"`
	Amount        float64 `json:"amount" binding:"required"`
}

func (server *Server) createTransaction(ctx *gin.Context) {
	var req createTransactionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	ac, err := server.store.GetAccount(ctx, req.AccountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	balance := ac.Balance
	var nb float64
	var tt string
	if req.OperationType == DEBIT {
		nb = balance - req.Amount
		tt = "debit"
	} else if req.OperationType == CREDIT {
		nb = balance - req.Amount
		tt = "credit"
	} else if req.OperationType == WITHDRAWL {
		nb = balance - req.Amount
		tt = "withdrawl"
	} else if req.OperationType == PAYMENT {
		nb = balance + req.Amount
		tt = "payment"
	}

	_, err = server.store.UpdateAccount(ctx, db.UpdateAccountParams{
		ID:      req.AccountID,
		Balance: nb,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	tr, err := server.store.CreateTransaction(ctx, db.CreateTransactionParams{
		AccountID:       req.AccountID,
		Amount:          req.Amount,
		TransactionType: tt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tr)
}
