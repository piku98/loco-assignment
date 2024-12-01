package dtos

type GetTransactionByIdDto struct {
	TransactionId string `uri:"transaction_id" binding:"required,numericString"`
}
