package usecase

import "amarthaTest/database"

func NewBillingEngineUsecase(billingEngineRP database.BillingEngine) BillingEngine {
	return &billingEngineUsecase{
		billingEngineRP: billingEngineRP,
	}
}

// on this part we will create logic to use logic on database that need to be integrate
