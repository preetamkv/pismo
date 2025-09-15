package accounts

import (
	"fmt"
)

// Model for the create account request body
type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

// Model for the create account response body
type CreateAccountResponse struct {
	AccountNumber string `json:"account_number"`
}

// Validate validates the body of create account request
func (r *CreateAccountRequest) Validate() error {
	if r.DocumentNumber == "" {
		return fmt.Errorf("document number is required")
	}

	return nil
}
