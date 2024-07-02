package database

import (
	"context"
	"database/sql"
	"time"

	"xorm.io/xorm"
)

func NewBillingEngineRepo(db *xorm.Engine) BillingEngine {
	return &billingEngineRepo{
		db: db,
	}
}

func (r *billingEngineRepo) GetOutstanding(ctx context.Context, loanID int) (float64, error) {
	var outstandingAmount float64

	sess := r.db.NewSession()
	defer sess.Close()

	err := sess.DB().QueryRow("SELECT OutstandingAmount FROM Loans WHERE LoanID = ?", loanID).Scan(&outstandingAmount)
	if err != nil {
		return 0, err
	}
	return outstandingAmount, nil
}

func (r *billingEngineRepo) IsDelinquent(ctx context.Context, loanID int) (bool, error) {
	sess := r.db.NewSession()
	defer sess.Close()

	rows, err := sess.DB().Query("SELECT Status FROM LoanSchedules WHERE LoanID = ? ORDER BY WeekNumber", loanID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	consecutiveMissedWeeks := 0
	for rows.Next() {
		var status string
		err := rows.Scan(&status)
		if err != nil {
			return false, err
		}
		if status == "Unpaid" {
			consecutiveMissedWeeks++
			if consecutiveMissedWeeks >= 2 {
				return true, nil
			}
		} else {
			consecutiveMissedWeeks = 0
		}
	}
	return false, nil
}

func (r *billingEngineRepo) MakePayment(ctx context.Context, loanID int, paymentAmount float64) (string, error) {
	sess := r.db.NewSession()
	defer sess.Close()

	// Check if paymentAmount is valid
	if paymentAmount != 110000 {
		return "Failure: Invalid payment amount", nil
	}

	// Begin a transaction
	err := sess.Begin()
	if err != nil {
		return "Failure", err
	}
	defer sess.Rollback()

	// Get the next unpaid schedule
	var scheduleID, weekNumber int
	var amountDue float64
	err = sess.DB().QueryRow("SELECT ScheduleID, WeekNumber, AmountDue FROM LoanSchedules WHERE LoanID = ? AND Status = 'Unpaid' ORDER BY WeekNumber LIMIT 1", loanID).Scan(&scheduleID, &weekNumber, &amountDue)
	if err != nil {
		if err == sql.ErrNoRows {
			return "Failure: No unpaid schedule found", nil
		}
		return "Failure", err
	}

	// Update the LoanSchedule
	_, err = sess.DB().Exec("UPDATE LoanSchedules SET AmountPaid = ?, Status = 'Paid' WHERE ScheduleID = ?", paymentAmount, scheduleID)
	if err != nil {
		return "Failure", err
	}

	// Update the Loan
	_, err = sess.DB().Exec("UPDATE Loans SET OutstandingAmount = OutstandingAmount - ? WHERE LoanID = ?", paymentAmount, loanID)
	if err != nil {
		return "Failure", err
	}

	// Insert the payment record
	_, err = sess.DB().Exec("INSERT INTO Payments (LoanID, PaymentDate, Amount) VALUES (?, ?, ?)", loanID, time.Now(), paymentAmount)
	if err != nil {
		return "Failure", err
	}

	// Commit the transaction
	err = sess.Commit()
	if err != nil {
		return "Failure", err
	}

	return "Success", nil
}
