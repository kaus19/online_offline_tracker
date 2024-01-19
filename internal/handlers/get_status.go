package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kaus19/online_offline_tracker/api"
	"github.com/kaus19/online_offline_tracker/internal/tools"

	"github.com/go-redis/redis"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

var ErrUnauthorized = errors.New("invalid username or token")
var database *redis.Client
var err error

func InitiateDb() {

	database, err = tools.NewDatabase()
	if err != nil {
		log.Error(err)
		return
	}
}

func GetUserStatus(w http.ResponseWriter, r *http.Request) {
	var params = api.StatusParams{}
	var decoder *schema.Decoder = schema.NewDecoder()

	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var userDetails *tools.UserDetails = tools.GetUserDetails(database, params.Username)
	if userDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.StatusResponse{
		Status: (*userDetails).Status,
		Code:   http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

func GetAllStatus(w http.ResponseWriter, r *http.Request) {
	var allUserDetails *[]tools.UserDetails = tools.GetAllStatus(database)
	if allUserDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var allResponses []api.UserDetailsResponse
	var allResponse api.UserDetailsResponse
	for _, v := range *allUserDetails {
		allResponse.Username = v.Username
		allResponse.Status = v.Status
		allResponses = append(allResponses, allResponse)
	}
	var response = api.AllStatusResponse{
		AllDetails: allResponses,
		Code:       http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
