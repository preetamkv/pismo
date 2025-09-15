package models

// Model for Account table
type Account struct {
	AccountID      string `json:"account_id" gorm:"primaryKey"`
	DocumentNumber string `json:"document_number" gorm:"unique"`
}

// Model for Transaction Table
type Transaction struct {
	TransactionID string  `json:"transaction_id" gorm:"primaryKey"`
	AccountID     string  `json:"account_id"`
	OperationType int     `json:"operation_type"`
	Amount        float64 `json:"amount"`
	EventDate     string  `json:"event_date"`
}
