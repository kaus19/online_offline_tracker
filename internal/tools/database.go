package tools

import log "github.com/sirupsen/logrus"

type LoginDetails struct{
	AuthToken string
	Username string
}

type StatusDetails struct{
	Status bool
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserStatus(username string) *StatusDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error){
	var database DatabaseInterface = &mockDB{}

	var err error = database.SetupDatabase()
	if err!=nil{
		log.Error(err)
		return nil, err
	}
	return &database, nil
}