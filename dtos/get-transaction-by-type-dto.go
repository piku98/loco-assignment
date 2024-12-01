package dtos

type GetTransactionByType struct {
	TransactionType string `uri:"type" binding:"required"`
}
