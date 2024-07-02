package usecase

import (
	"amarthaTest/database"
)

type (
	billingEngineUsecase struct {
		billingEngineRP database.BillingEngine
	}

	BillingEngine interface {
	}
)
