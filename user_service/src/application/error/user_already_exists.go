package error

type UserAlreadyExistsError struct {
	Email string
}

func (r *UserAlreadyExistsError) Error() string {
	return "User already exists with email: " + r.Email
}
