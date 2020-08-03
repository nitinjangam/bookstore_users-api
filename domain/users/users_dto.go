package users

import (
	"strings"

	"github.com/nitinjangam/bookstore_users-api/utils/errors"
)

const (
	//StatusActive variable to store user as active user
	StatusActive = "active"
)

//User struct
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

//Users typedef
type Users []User

//Validate function
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.CreateNewBadReqError("Invalid Email Address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.CreateNewBadReqError("Invalid Password")
	}
	return nil
}
