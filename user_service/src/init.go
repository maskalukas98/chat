package src

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"user_service/src/application/use_case"
	"user_service/src/common/database"
	"user_service/src/domain/input_port"
	"user_service/src/domain/output_port"
	"user_service/src/infrastructure/messaging"
	"user_service/src/infrastructure/repository"
)

var (
	_             = godotenv.Load()
	mainDbHost    = os.Getenv("MAIN_DB_HOST")
	mainDbPortStr = os.Getenv("MAIN_DB_PORT")
	mainDbPort, _ = strconv.Atoi(mainDbPortStr)
	mainDbUser    = os.Getenv("MAIN_DB_USER")
	mainDbPass    = os.Getenv("MAIN_DB_PASS")
	mainDbDbName  = os.Getenv("MAIN_DB_DBNAME")

	kafkaUrl = os.Getenv("KAFKA_URL")

	userSqlDb, _ = database.NewPostgresUserDatabase(database.DbConfig{
		Host:     mainDbHost,
		Port:     mainDbPort,
		User:     mainDbUser,
		Password: mainDbPass,
		DBName:   mainDbDbName,
		SSLMode:  "disable",
	})

	postgresUserRepository output_port.UserRepositoryPort = repository.NewUserRepository(userSqlDb)

	kafkaProducerFactory = messaging.KafkaProducerFactory{}

	kafkaProducer output_port.MessagingProducer = messaging.NewKafkaProducer(
		kafkaProducerFactory.CreateSyncProducer(kafkaUrl),
	)

	CreateUserUseCase input_port.CreateUserPort = use_case.NewCreateUserUseCase(&postgresUserRepository, &kafkaProducer)
	GetUserUseCase    input_port.GetUserPort    = use_case.NewGetUserUseCase(&postgresUserRepository)
)

func init() {
	log.Println("Global initialization completed.")
}
