package dto

import "time"

type LoanSchedule struct {
	ScheduleID int
	LoanID     int
	WeekNumber int
	DueDate    time.Time
	AmountDue  float64
	AmountPaid float64
	Status     string
}
