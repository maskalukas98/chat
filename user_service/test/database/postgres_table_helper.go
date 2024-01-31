package test_database

import (
	"database/sql"
	"user_service/src/domain"
	"user_service/src/domain/output_port"
	"user_service/src/domain/value_object"
	"user_service/src/infrastructure/repository"
)

func ConvertRowsToUsers(rows *sql.Rows) []domain.User {
	var users []domain.User

	for rows.Next() {
		var user domain.User
		var userId int64
		rows.Scan(&userId, &user.Name, &user.Email)
		user.UserId = *value_object.NewUserId(userId)

		users = append(users, user)
	}

	return users
}

func BeforeEach() (*PostgresDb, output_port.UserRepositoryPort) {
	postgres := NewPostgresDb()
	userRepository := repository.NewUserRepository(postgres.GetDb())

	postgres.ClearTable("users")

	return postgres, userRepository
}

func AfterEach(postgres *PostgresDb) {
	postgres.GetDb().Close()
}
