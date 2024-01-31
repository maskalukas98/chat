package port_output

import _type "chat_service/src/domain/type"

type UserService interface {
	GetUserById(userId int64) *_type.User
}
