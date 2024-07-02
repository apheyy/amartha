package controller

import (
	"amarthaTest/config"
	"amarthaTest/usecase"
	"log"
	"net/http"
)

type billingEngineAPI struct {
	billingEngineUC usecase.BillingEngine
}

func NewBillingEngineAPI(billingEngineUC usecase.BillingEngine) *billingEngineAPI {
	return &billingEngineAPI{
		billingEngineUC: billingEngineUC,
	}
}

func (billingEngineAPI *billingEngineAPI) TestingApi(w http.ResponseWriter, r *http.Request) {
	// ctx := context.Background()
	funcName := "TestingApi"
	log.Printf("[%s] Start Testing API", funcName)

	resp := "TestingApi"

	config.ResponseJSON(w, resp)
}

// on this part we will integrate end to end to use all the logic api on usecase
