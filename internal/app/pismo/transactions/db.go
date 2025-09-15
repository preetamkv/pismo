package transactions

import (
	"time"

	"github.com/google/uuid"
	"github.com/preetamkv/pismo/internal/pkg/models"

	"gorm.io/gorm"
)

// createTransaction creates a new Transaction entry in the DB
func createTransaction(db *gorm.DB, req *CreateTransactionRequest) (string, error) {
	txID := uuid.New() // Generate transaction ID

	amt := req.Amount
	if req.OperationType <= 3 {
		amt = req.Amount * -1 // Ensure amount is changed to negative based on operation type
	}

	// Store transaction time in required format
	txTime := time.Now().UTC().Format("2006-01-02T15:04:05.0000000Z")

	tx := models.Transaction{
		TransactionID: txID.String(),
		AccountID:     req.AccountID,
		OperationType: req.OperationType,
		Amount:        amt,
		EventDate:     txTime,
	}

	// Store the transaction in DB
	err := db.Create(tx).Error
	if err != nil {
		return "", err
	}
	return tx.TransactionID, nil
}
