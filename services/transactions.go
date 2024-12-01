package services

import (
	"errors"
	"fmt"
	customerrors "loco-assignment/custom-errors"
	"loco-assignment/db"
	"loco-assignment/dtos"
	"loco-assignment/models"

	"gorm.io/gorm"
)

type TransactionService struct {
	postgresClient *db.PostgresClient
}

func NewTrasactionService(postgresClient *db.PostgresClient) *TransactionService {
	return &TransactionService{postgresClient: postgresClient}
}

func (service *TransactionService) UpsertTransaction(
	transactionId int64,
	inputDetails dtos.PutTransactionBodyDto,
) error {
	client := service.postgresClient.GetClient()

	if inputDetails.ParentID != nil {
		pp, err := service.GetTransactionById(*(inputDetails.ParentID))
		fmt.Println(pp, err)
		if err != nil {
			if _, ok := err.(customerrors.NotFoundError); ok {
				return customerrors.NotFoundError{
					Message: fmt.Sprintf("Parent with transaction id %d not found", *(inputDetails.ParentID)),
				}
			}
			return err
		}
	}

	transaction := models.Transaction{
		ID:     transactionId,
		Amount: inputDetails.Amount,
		Type:   inputDetails.Type,
	}

	if inputDetails.ParentID != nil {
		transaction.ParentID = *inputDetails.ParentID
	}

	err := client.Save(&transaction).Error
	return err
}

func (service *TransactionService) GetTransactionById(
	transactionId int64,
) (models.Transaction, error) {
	client := service.postgresClient.GetClient()
	var transaction models.Transaction
	if err := client.Where("id = ?", transactionId).First(&transaction).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Transaction{},
				customerrors.NotFoundError{
					Message: fmt.Sprintf("Transaction with id %d not found", transactionId),
				}
		}
		return models.Transaction{}, err
	}
	return transaction, nil
}

func (service *TransactionService) GetTransactionByType(
	transactionType string,
) ([]int64, error) {
	client := service.postgresClient.GetClient()
	var transactionIds []int64

	err := client.Model(&models.Transaction{}).Where(
		"type = ?", transactionType,
	).Pluck("id", &transactionIds).Error

	return transactionIds, err
}

func (service *TransactionService) TransactionSum(
	transactionId int64,
) (float64, error) {
	client := service.postgresClient.GetClient()
	var sum float64
	err := client.Raw(`
	WITH RECURSIVE transaction_tree AS (
			SELECT id, amount, parent_id
			FROM transactions
			WHERE id = ?
			UNION ALL
			SELECT t.id, t.amount, t.parent_id
			FROM transactions t
			INNER JOIN transaction_tree tt ON t.parent_id = tt.id
		)
		SELECT SUM(amount) AS total_amount FROM transaction_tree;
	`, transactionId).Scan(&sum).Error

	return sum, err
}
