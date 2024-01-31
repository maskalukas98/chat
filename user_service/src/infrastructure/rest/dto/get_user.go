package dto

// GetUserResponseDto represents the DTO for user response.
type GetUserResponseDto struct {
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
