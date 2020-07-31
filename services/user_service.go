package services

import (
	"github.com/nitinjangam/bookstore_users-api/domain/users"
	"github.com/nitinjangam/bookstore_users-api/utils/errors"
)

//CreateUser function
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
