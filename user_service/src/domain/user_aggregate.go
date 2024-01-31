package domain

import "user_service/src/domain/value_object"

type User struct {
	UserId value_object.UserId
	Name   string
	Email  string
}
