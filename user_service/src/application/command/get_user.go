package command

type GetUserByIdRequest struct {
	UserId int64
}

type GetUserResponse struct {
	UserId int64
	Name   string
	Email  string
}
