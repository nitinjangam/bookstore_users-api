package users

import (
	"encoding/json"
)

//PublicUser struct for public requests
type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

//PrivateUser struct for private requests
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

//Marshall function
func (user *User) Marshall(isPublic bool) interface{} {
	userJSON, _ := json.Marshal(user)
	if isPublic {
		var publicUser PublicUser
		json.Unmarshal(userJSON, &publicUser)
		return publicUser
	}
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}

//Marshall function on slice of user
func (users Users) Marshall(isPublic bool) []interface{} {

	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}
