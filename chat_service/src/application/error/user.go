package custom_error

import "fmt"

type UserNotFoundError struct {
	UserId int64
}

func (r *UserNotFoundError) Error() string {
	return fmt.Sprintf("User not found: %d", r.UserId)
}
