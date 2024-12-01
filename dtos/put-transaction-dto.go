package dtos

type PutTransactionBodyDto struct {
	Amount   float64 `json:"amount" binding:"required"`
	Type     string  `json:"type" binding:"required"`
	ParentID *int64  `json:"parent_id"`
}

type PutTransactionParamsDto struct {
	TransactionId string `uri:"transaction_id" binding:"required,numericString"`
}
