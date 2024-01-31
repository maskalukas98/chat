package test_repository

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	custom_error "user_service/src/application/error"
	"user_service/src/domain/output_port"
	"user_service/src/domain/value_object"
	test_database "user_service/test/database"
)

var postgres *test_database.PostgresDb
var userRepository output_port.UserRepositoryPort

func TestPostgresUserRepository_CreateUser_Success(t *testing.T) {
	// prepare
	postgres, userRepository = test_database.BeforeEach()

	// execute
	userId, err := userRepository.CreateUser("lukas", "lukas.maska@fakeEmail.com")

	// validation
	rows := postgres.GetAll("users")
	users := test_database.ConvertRowsToUsers(rows)

	assert.Equal(t, len(users), 1)
	assert.Equal(t, users[0].UserId.GetId(), int64(1))
	assert.Equal(t, users[0].Name, "lukas")
	assert.Equal(t, users[0].Email, "lukas.maska@fakeEmail.com")
	assert.Equal(t, userId.GetId(), int64(1))
	assert.Nil(t, err)

	test_database.AfterEach(postgres)
}

func TestPostgresUserRepository_CreateTwoUsers_Success(t *testing.T) {
	// prepare
	postgres, userRepository = test_database.BeforeEach()

	// execute
	userId, err := userRepository.CreateUser("josef", "josef.novotny@fakeEmail.com")
	userId2, err2 := userRepository.CreateUser("lukas", "lukas.maska@fakeEmail.com")

	// validation
	rows := postgres.GetAll("users")
	users := test_database.ConvertRowsToUsers(rows)

	assert.Equal(t, len(users), 2)

	assert.Equal(t, users[0].UserId.GetId(), int64(1))
	assert.Equal(t, users[0].Name, "josef")
	assert.Equal(t, users[0].Email, "josef.novotny@fakeEmail.com")
	assert.Equal(t, userId.GetId(), int64(1))
	assert.Nil(t, err)

	assert.Equal(t, users[1].UserId.GetId(), int64(2))
	assert.Equal(t, users[1].Name, "lukas")
	assert.Equal(t, users[1].Email, "lukas.maska@fakeEmail.com")
	assert.Equal(t, userId2.GetId(), int64(2))
	assert.Nil(t, err2)

	test_database.AfterEach(postgres)
}

func TestPostgresUserRepository_CreateUser_Fail_DuplicateEmail(t *testing.T) {
	// prepare
	postgres, userRepository = test_database.BeforeEach()

	// execute
	userId, err := userRepository.CreateUser("josef", "lukas.maska@fakeEmail.com")
	userId2, err := userRepository.CreateUser("lukas", "lukas.maska@fakeEmail.com")

	// validation
	rows := postgres.GetAll("users")
	users := test_database.ConvertRowsToUsers(rows)

	assert.Equal(t, len(users), 1)
	assert.Equal(t, users[0].UserId.GetId(), int64(1))
	assert.Equal(t, users[0].Name, "josef")
	assert.Equal(t, users[0].Email, "lukas.maska@fakeEmail.com")
	assert.Equal(t, userId.GetId(), int64(1))

	var userExistsErr *custom_error.UserAlreadyExistsError
	assert.True(t, errors.As(err, &userExistsErr))

	assert.Nil(t, userId2)

	test_database.AfterEach(postgres)
}

func TestPostgresUserRepository_GetUserById_Success(t *testing.T) {
	// prepare
	postgres, userRepository = test_database.BeforeEach()

	userRepository.CreateUser("lukas", "lukas.maska@fakeEmail.com")
	userRepository.CreateUser("josef", "josef.novotny@fakeEmail.com")

	// execute
	user, err := userRepository.GetUserById(*value_object.NewUserId(2))

	// validation
	assert.Equal(t, user.UserId.GetId(), int64(2))
	assert.Equal(t, user.Name, "josef")
	assert.Equal(t, user.Email, "josef.novotny@fakeEmail.com")
	assert.Nil(t, err)

	test_database.AfterEach(postgres)
}

func TestPostgresUserRepository_GetUserById_NotFound(t *testing.T) {
	// prepare
	postgres, userRepository = test_database.BeforeEach()

	// execute
	user, err := userRepository.GetUserById(*value_object.NewUserId(10))

	// validation
	assert.Nil(t, user)
	assert.Nil(t, err)

	test_database.AfterEach(postgres)
}
