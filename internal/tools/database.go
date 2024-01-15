package tools

import (
	log "github.com/sirupsen/logrus"
	"github.com/go-redis/redis"
	"encoding/json"
)

type UserDetails struct{
	AuthToken string `json:"token"`
	Status bool `json:"status"`
}

func NewDatabase() (*redis.Client, error){

	client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })

	// Adding Initial Users
	user1, err := json.Marshal(UserDetails{AuthToken: "12345", Status: false})
    if err != nil {
        log.Error(err)
    }
	user2, err := json.Marshal(UserDetails{AuthToken: "123", Status: true})
    if err != nil {
        log.Error(err)
    }
	user3, err := json.Marshal(UserDetails{AuthToken: "1", Status: true})
    if err != nil {
        log.Error(err)
    }

	err = SetupDatabase(client, map[string]interface{}{
        "maloo": user1,
        "shashi": user2,
        "rishav": user3,
    })
	if err != nil {
		log.Error(err)
	}
	
	return client, nil
}

func SetupDatabase(client *redis.Client, values map[string]interface{}) error {
	for key, value := range values {
        err := client.Set(key, value, 0).Err() // 0 means no expiration
        if err != nil {
            return err
        }
    }
    return nil
}

func GetUserDetails(client *redis.Client, value string) *UserDetails {
 
	val, err := client.Get(value).Result()
	if err != nil {
		log.Error(err)
		return nil
	}

	var userDetails = UserDetails{}	
    err = json.Unmarshal([]byte(val), &userDetails)
    if err != nil {
        log.Error("Error unmarshalling JSON:", err)
        return nil
    }

	log.Info(userDetails)
	return &userDetails
}