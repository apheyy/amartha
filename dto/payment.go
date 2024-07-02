package dto

import "time"

type Payment struct {
	PaymentID   int
	LoanID      int
	PaymentDate time.Time
	Amount      float64
}
