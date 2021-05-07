package services

import (
	"github.com/mochanpypa1709/bookstore_users-api/domain/users"
	"github.com/mochanpypa1709/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
