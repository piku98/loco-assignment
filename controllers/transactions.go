package controllers

import (
	"fmt"
	customerrors "loco-assignment/custom-errors"
	"loco-assignment/dtos"
	"loco-assignment/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service *services.TransactionService
}

func NewTrasactionController(transactionService *services.TransactionService) *TransactionController {
	return &TransactionController{service: transactionService}
}

func (controller *TransactionController) RegisterRoutes(router *gin.Engine) {
	router.PUT("/transactionservice/transaction/:transaction_id", controller.UpsertTransaction)
	router.GET("/transactionservice/transaction/:transaction_id", controller.GetTransactionById)
	router.GET("/transactionservice/types/:type", controller.GetTransactionByType)
	router.GET("/transactionservice/sum/:transaction_id", controller.TransactionSum)
}

func (controller *TransactionController) UpsertTransaction(ctx *gin.Context) {
	var params dtos.PutTransactionParamsDto
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body dtos.PutTransactionBodyDto
	if err := ctx.ShouldBindBodyWithJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionId, _ := strconv.ParseInt(params.TransactionId, 10, 64)

	err := controller.service.UpsertTransaction(
		transactionId,
		body,
	)

	if err != nil {
		fmt.Printf("Error caused during execution %v", err)
		if _, ok := err.(customerrors.NotFoundError); ok {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (controller *TransactionController) GetTransactionById(ctx *gin.Context) {
	var params dtos.GetTransactionByIdDto

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionId, _ := strconv.ParseInt(params.TransactionId, 10, 64)

	transaction, err := controller.service.GetTransactionById(transactionId)
	if err != nil {
		status := http.StatusInternalServerError
		if _, ok := err.(customerrors.NotFoundError); ok {
			status = http.StatusNotFound
		}
		ctx.JSON(status, gin.H{"error": err.Error()})
		return
	}
	// Should implement a formatting function instead.
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"amount":    transaction.Amount,
			"type":      transaction.Type,
			"parent_id": transaction.ParentID,
		},
	)
}

func (controller *TransactionController) GetTransactionByType(ctx *gin.Context) {
	var params dtos.GetTransactionByType

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionIds, err := controller.service.GetTransactionByType(params.TransactionType)
	if err != nil {
		fmt.Printf("Get Transaction by type error %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong."})
		return
	}

	ctx.JSON(http.StatusOK, transactionIds)
}

func (controller *TransactionController) TransactionSum(ctx *gin.Context) {
	var params dtos.TransactionSumDto

	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transactionId, _ := strconv.ParseInt(params.TransactionId, 10, 64)

	sum, err := controller.service.TransactionSum(transactionId)
	if err != nil {
		fmt.Printf("Get Transaction sum error %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"sum": sum})
}
