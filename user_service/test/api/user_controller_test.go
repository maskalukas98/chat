package test_api

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"user_service/src/infrastructure/rest/dto"
	test_database "user_service/test/database"
)

// todo: create tests for createUser - duplicate email, getUserById

var postgres *test_database.PostgresDb

func TestUserController_CreateUser(t *testing.T) {
	// prepare
	postgres, _ = test_database.BeforeEach()

	t.Run("CreateUser", func(t *testing.T) {
		// execute
		payload := []byte(`{"name": "john", "email": "john@example.com"}`)
		response, err := http.Post("http://localhost:8080/api/v1/users", "application/json", bytes.NewBuffer(payload))
		assert.NoError(t, err)
		defer response.Body.Close()

		// validation
		assert.Equal(t, http.StatusCreated, response.StatusCode)

		var responseBody dto.CreateUserResponseDto
		json.NewDecoder(response.Body).Decode(&responseBody)
		assert.Equal(t, int64(1), responseBody.UserId)

		rows := postgres.GetAll("users")
		users := test_database.ConvertRowsToUsers(rows)
		assert.Equal(t, 1, len(users))
		assert.Equal(t, int64(1), users[0].UserId.GetId())
		assert.Equal(t, "john", users[0].Name)
		assert.Equal(t, "john@example.com", users[0].Email)
	})

	test_database.AfterEach(postgres)
}
