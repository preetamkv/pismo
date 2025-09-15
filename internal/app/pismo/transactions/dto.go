package transactions

import (
	"fmt"
)

// Model for the create transaction request body
type CreateTransactionRequest struct {
	AccountID     string  `json:"account_id"`
	OperationType int     `json:"operation_type_id"`
	Amount        float64 `json:"amount"`
}

type CreateTransactionResponse struct {
	TransactionID string `json:"transaction_id"`
}

// Validate validates the body of create transaction request
func (r *CreateTransactionRequest) Validate() error {
	if r.AccountID == "" {
		return fmt.Errorf("account ID is required")
	}
	if r.OperationType <= 0 || r.OperationType > 4 {
		return fmt.Errorf("invalid operation")
	}
	if r.Amount < 0 {
		return fmt.Errorf("amount is required to be positive")
	}
	/*
		Expecting amount to be positive in the request from the client,
		If the expectation is wrong, one can update validation here.
		It will have multiple conditions based on OpertionType value.
	*/

	return nil
}
