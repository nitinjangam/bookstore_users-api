package users

import "github.com/nitinjangam/bookstore_users-api/utils/errors"

//User struct
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

//Validate function
func (user *User) Validate() *errors.RestErr {
	if user.Email == "" {
		return errors.CreateNewBadReqError("Invalid Email Address")
	}
	return nil
}
