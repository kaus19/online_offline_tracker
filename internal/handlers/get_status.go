package handlers

import (
	"errors"
	"encoding/json"
	"net/http"
	"github.com/kaus19/online_offline_tracker/api"
	"github.com/kaus19/online_offline_tracker/internal/tools"

	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
	"github.com/go-redis/redis"
)

var ErrUnauthorized = errors.New("invalid username or token")

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

	var database *redis.Client
	database, err = tools.NewDatabase()
	if err!=nil{
		api.InternalErrorHandler(w)
		return
	}

	var userDetails *tools.UserDetails = tools.GetUserDetails(database, params.Username)
	if userDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
	var token = r.Header.Get("Authorization")
	if token != (*userDetails).AuthToken {
		log.Error(ErrUnauthorized)
		api.RequestErrorHandler(w, ErrUnauthorized)
		return
	}

	var response = api.StatusResponse{
		Status: (*userDetails).Status,
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