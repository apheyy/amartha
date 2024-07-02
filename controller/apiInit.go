package controller

import (
	"amarthaTest/config"
	"amarthaTest/database"
	"amarthaTest/model"
	"amarthaTest/usecase"
)

func NewAPI(
	billingEngineUC usecase.BillingEngine,
) *APIModule {
	return &APIModule{
		billingEngineAPI: NewBillingEngineAPI(billingEngineUC),
	}
}

type Usecase struct {
	BillingEngineUC usecase.BillingEngine
}

func InitUsecase(cfg model.AppConfig) (uc Usecase) {
	db := config.InitDB(cfg)

	billingEngineRP := database.NewBillingEngineRepo(db)

	uc.BillingEngineUC = usecase.NewBillingEngineUsecase(billingEngineRP)

	return uc
}
