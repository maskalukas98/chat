package command

type CreateUserRequest struct {
	Name      string
	Email     string
	IpAddress string
}

type CreateUserResponse struct {
	UserId int64
}
