package src

import (
	"chat_service/src/application/use_case"
	"chat_service/src/common/database"
	"chat_service/src/common/redis"
	port_input "chat_service/src/domain/port/input"
	port_output "chat_service/src/domain/port/output"
	"chat_service/src/infrastructure/repository"
	service_rest "chat_service/src/infrastructure/service/rest"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	_            = godotenv.Load()
	mainDbUser   = os.Getenv("MAIN_DB_USER")
	mainDbPass   = os.Getenv("MAIN_DB_PASS")
	mainDbHost   = os.Getenv("MAIN_DB_HOST")
	mainDbPort   = os.Getenv("MAIN_DB_PORT")
	mainDbDbName = os.Getenv("MAIN_DB_DBNAME")

	sentinelHost       = os.Getenv("SENTINEL_HOST")
	sentinelPort       = os.Getenv("SENTINEL_PORT")
	sentinelMasterName = os.Getenv("SENTINEL_MASTER_NAME")

	userServiceUrl = os.Getenv("USER_SERVICE_URL")

	chatMongoDb, _ = database.NewMongoDBClient(
		"mongodb://" + mainDbUser + ":" + mainDbPass + "@" + mainDbHost + ":" + mainDbPort + "/" + mainDbDbName,
	)

	sentinel = redis.NewSentinel(sentinelMasterName, sentinelHost+":"+sentinelPort)

	mongodbMessageRepository port_output.MessageRepository = repository.NewMongodbMessageRepository(
		chatMongoDb,
	)

	redisRecentMessagesCacheRepository port_output.LatestMessagesCacheRepository = repository.NewRedisRecentMessagesRepository(
		sentinel.GetRedisClient(),
	)

	restUserService port_output.UserService = service_rest.NewRestUserService(userServiceUrl)

	SendMessageUseCase port_input.SendMessage = use_case.NewSendMessageUseCase(
		&mongodbMessageRepository,
		&redisRecentMessagesCacheRepository,
		&restUserService,
	)

	GetMessagesUseCase port_input.GetMessages = use_case.NewGetMessagesUseCase(
		&mongodbMessageRepository,
		&redisRecentMessagesCacheRepository,
		&restUserService,
	)
)

func init() {
	log.Println("Global initialization completed.")
}
