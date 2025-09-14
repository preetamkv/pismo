package accounts

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/preetamkv/pismo/internal/pkg/models"
	"gorm.io/gorm"
)

func createAccount(db *gorm.DB, req *CreateAccountRequest) (string, error) {
	accID := uuid.New()
	acc := models.Account{
		AccountID:      accID.String(),
		DocumentNumber: req.DocumentNumber,
	}
	if err := db.Create(acc).Error; err != nil {
		return "", fmt.Errorf("failed to create Account")
	}
	return acc.AccountID, nil
}

func fetchAccount(db *gorm.DB, accID string) (*models.Account, error) {
	var acc models.Account
	result := db.First(&acc, accID)
	if result.Error != nil {
		return &models.Account{}, fmt.Errorf("unable to fetch Account")
	}
	return &acc, nil
}
