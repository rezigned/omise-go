package omise

import (
	"time"
)

// Transfer represents Transfer object.
// See https://www.omise.co/transfers-api for more information.
type Transfer struct {
	Base
	Amount int `json:"amount"`
	BankAccount *BankAccount `json:"bank_account"`
	Currency string `json:"currency"`
	FailFast bool `json:"fail_fast"`
	FailureCode *string `json:"failure_code"`
	FailureMessage *string `json:"failure_message"`
	Fee int `json:"fee"`
	FeeVat int `json:"fee_vat"`
	Location string `json:"location"`
	Metadata map[string]interface{} `json:"metadata"`
	Net int `json:"net"`
	Paid bool `json:"paid"`
	PaidAt *time.Time `json:"paid_at"`
	Recipient string `json:"recipient"`
	Schedule string `json:"schedule"`
	Sendable bool `json:"sendable"`
	Sent bool `json:"sent"`
	SentAt *time.Time `json:"sent_at"`
	TotalFee int `json:"total_fee"`
	Transactions []Transaction `json:"transactions"`
}