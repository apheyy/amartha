package config

import (
	"amarthaTest/common/constants"
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", constants.APP_JSON)
	json.NewEncoder(w).Encode(resp)
}
