package transaction

import "time"

type TransactionPayload struct {
	Debit           float64                `json:"debit"`
	DebitType       string                 `json:"debit_type"`
	Credit          float64                `json:"credit"`
	CreditType      string                 `json:"credit_type"`
	TransactionType string                 `json:"transaction_type"`
	TransactionTime time.Time              `json:"transaction_time"`
	Extra           map[string]interface{} `json:"extra"`
}
