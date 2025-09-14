package accounts

import (
	"fmt"
)

type CreateAccountRequest struct {
	DocumentNumber string `json:"document_number"`
}

func (r *CreateAccountRequest) Validate() error {
	if r.DocumentNumber == "" {
		return fmt.Errorf("document number is required")
	}

	return nil
}
