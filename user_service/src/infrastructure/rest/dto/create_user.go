package dto

type CreateUserRequestDto struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserResponseDto struct {
	UserId int64             `json:"user_id"`
	Links  map[string]string `json:"links"`
}
