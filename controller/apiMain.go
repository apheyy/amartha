package controller

import (
	"amarthaTest/common/constants"
	"amarthaTest/config"
	"amarthaTest/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func HandleAPI(cfg model.AppConfig) {
	muxRouter := mux.NewRouter().StrictSlash(true)
	r := muxRouter.PathPrefix(constants.AMARTHA).Subrouter()

	ng := negroni.New()
	ng.Use(negroni.HandlerFunc(config.CaptureNegroniHandler))

	uc := InitUsecase(cfg)
	apiModule := NewAPI(uc.BillingEngineUC)

	r.HandleFunc(constants.V1, apiModule.billingEngineAPI.TestingApi).Methods("POST")

	// on this part we will integrate api that will use by our system

	log.Println("Loading API routes ....")
	ng.UseHandler(muxRouter)

	apiPort := 8000
	log.Println("api running on port:", apiPort)
	http.ListenAndServe(":"+strconv.Itoa(apiPort), ng)
}
