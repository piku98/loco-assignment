package dtos

type TransactionSumDto struct {
	TransactionId string `uri:"transaction_id" binding:"required,numericString"`
}
