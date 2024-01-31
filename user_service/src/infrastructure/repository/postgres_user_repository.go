package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	custom_error "user_service/src/application/error"
	"user_service/src/domain"
	"user_service/src/domain/value_object"

	_ "github.com/lib/pq"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) CreateUser(username, email string) (*value_object.UserId, error) {
	var userID int64
	err := ur.DB.QueryRow(
		"INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		username, email,
	).Scan(&userID)

	if err != nil {
		pqErr, isPqError := err.(*pq.Error)
		if isPqError && pqErr.Code == "23505" {
			return nil, &custom_error.UserAlreadyExistsError{}
		}

		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	return value_object.NewUserId(userID), nil
}

func (r *UserRepository) GetUserById(userId value_object.UserId) (*domain.User, error) {
	query := "SELECT name, email FROM users WHERE id = $1 LIMIT 1"
	row := r.DB.QueryRow(query, userId.GetId())

	var user domain.User

	err := row.Scan(&user.Name, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	user.UserId = userId

	return &user, nil
}
