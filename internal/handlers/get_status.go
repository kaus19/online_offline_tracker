package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/kaus19/online_offline_tracker/api"
	"github.com/kaus19/online_offline_tracker/internal/tools"

	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetUserStatus(w http.ResponseWriter, r *http.Request) {
	var params = api.StatusParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())
	if err!=nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err!=nil{
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.StatusDetails = (*database).GetUserStatus(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.StatusResponse{
		Status: (*tokenDetails).Status,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err!=nil{
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}