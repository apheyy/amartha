package database

import (
	"context"

	"xorm.io/xorm"
)

type (
	billingEngineRepo struct {
		db *xorm.Engine
	}

	BillingEngine interface {
		GetOutstanding(ctx context.Context, loanID int) (float64, error)
		IsDelinquent(ctx context.Context, loanID int) (bool, error)
		MakePayment(ctx context.Context, loanID int, paymentAmount float64) (string, error)
	}
)
