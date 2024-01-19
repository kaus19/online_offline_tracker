package tools

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

const MAX_ENTRIES = 10

type UserDetails struct {
	Username string `json:"username"`
	Status   bool   `json:"status"`
}

func NewDatabase() (*redis.Client, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Adding Initial Users
	user1, err := json.Marshal(UserDetails{Username: "maloo", Status: true})
	if err != nil {
		log.Error(err)
	}
	user2, err := json.Marshal(UserDetails{Username: "sashi", Status: true})
	if err != nil {
		log.Error(err)
	}
	user3, err := json.Marshal(UserDetails{Username: "rishav", Status: true})
	if err != nil {
		log.Error(err)
	}

	err = SetupDatabase(client, map[string]interface{}{
		"maloo":  user1,
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
		err := client.Set(key, value, 30*time.Second).Err() //30s ttl
		if err != nil {
			return err
		}
	}
	return nil
}

func AddUser(client *redis.Client, value string) error {
	userDetail, err := json.Marshal(UserDetails{Username: value, Status: true})
	if err != nil {
		log.Error(err)
	}

	err = client.Set(value, userDetail, 30*time.Second).Err() //30s ttl
	if err != nil {
		return err
	}
	return nil
}

func GetUserDetails(client *redis.Client, value string) *UserDetails {

	val, err := client.Get(value).Result()
	if err != nil {
		if err == redis.Nil {
			err = AddUser(client, value)
			if err != nil {
				log.Error(err)
				return nil
			}
			return &UserDetails{Username: value, Status: true}
		}
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

func GetAllStatus(client *redis.Client) *[]UserDetails {

	var cursor uint64
	var n int
	var allStatus = []UserDetails{}
	var userDetails = UserDetails{}

	for {
		// Scan with the current cursor position
		var keys []string
		var err error
		keys, cursor, err = client.Scan(cursor, "*", 10).Result()
		if err != nil {
			log.Error(err)
		}

		// Process the batch of keys
		for _, key := range keys {
			value, err := client.Get(key).Result()
			if err != nil {
				log.Printf("Error getting value for key %s: %v\n", key, err)
				continue
			}
			log.Printf("Key: %s, Value: %s\n", key, value)
			err = json.Unmarshal([]byte(value), &userDetails)
			if err != nil {
				log.Error("Error unmarshalling JSON:", err)
				return nil
			}
			allStatus = append(allStatus, userDetails)
		}

		// Break if we have completed a full iteration
		if cursor == 0 {
			break
		}
		n++
		if n > MAX_ENTRIES {
			break
		}
	}
	return &allStatus
}
