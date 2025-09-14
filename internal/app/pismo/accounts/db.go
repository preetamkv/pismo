package accounts

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/preetamkv/pismo/internal/pkg/models"
	"gorm.io/gorm"
)

// createAccount creates a new Account entry in the DB
func createAccount(db *gorm.DB, req *CreateAccountRequest) (string, error) {
	accID := uuid.New() // Generate Acc ID
	acc := models.Account{
		AccountID:      accID.String(),
		DocumentNumber: req.DocumentNumber,
	}

	// Store the data in DB
	err := db.Create(acc).Error
	if err != nil {
		return "", fmt.Errorf("failed to create Account")
	}
	return acc.AccountID, nil
}

// fetchAccount fetches the first matching account's data from the DB based on the given account ID
func fetchAccount(db *gorm.DB, accID string) (*models.Account, error) {
	var acc models.Account
	result := db.First(&acc, accID)
	if result.Error != nil {
		return &models.Account{}, fmt.Errorf("unable to fetch Account")
	}
	return &acc, nil
}
