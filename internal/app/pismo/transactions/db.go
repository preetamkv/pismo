package transactions

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/preetamkv/pismo/internal/pkg/models"
	"gorm.io/gorm"
)

func createTransaction(db *gorm.DB, req *CreateTransactionRequest) (string, error) {
	txID := uuid.New()
	// Add check if the account exists.
	var amt int64
	if req.OperationType <= 3 {
		amt = req.Amount * -1
	}
	txTime := time.Now().UTC().Format("2006-01-02T15:04:05.0000000Z")
	tx := models.Transaction{
		TransactionID: txID.String(),
		AccountID:     req.AccountID,
		OperationType: req.OperationType,
		Amount:        amt,
		EventDate:     txTime,
	}
	if err := db.Create(tx).Error; err != nil {
		return "", fmt.Errorf("failed to create Transaction")
	}
	return tx.TransactionID, nil
}
