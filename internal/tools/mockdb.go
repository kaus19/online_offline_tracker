package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "12345",
		Username: "alex",
	},
}

var mockStatusDetails = map[string]StatusDetails{
	"alex": {
		Status: true,
		Username: "alex",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	//simulate db call
	time.Sleep(time.Second *1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserStatus(username string) *StatusDetails {
	//simulate db call
	time.Sleep(time.Second *1)
	var clientData = StatusDetails{}
	clientData, ok := mockStatusDetails[username]

	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}