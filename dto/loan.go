package dto

import "time"

type Loan struct {
	LoanID            int
	CustomerID        int
	PrincipalAmount   float64
	InterestRate      float64
	TotalAmount       float64
	StartDate         time.Time
	EndDate           time.Time
	OutstandingAmount float64
}
