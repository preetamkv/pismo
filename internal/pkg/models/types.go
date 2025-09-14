package models

type Account struct {
	AccountID      string `json:"account_id" gorm:"primaryKey"`
	DocumentNumber string `json:"document_number" gorm:"unique"`
}

type Transaction struct {
	TransactionID string `json:"transaction_id" gorm:"primaryKey"`
	AccountID     string `json:"account_id"`
	OperationType uint8  `json:"operation_type"`
	Amount        int64  `json:"amount"`
	EventDate     string `json:"event_date"`
}
