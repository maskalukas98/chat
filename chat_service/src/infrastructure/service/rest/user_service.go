package service_rest

import (
	_type "chat_service/src/domain/type"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type RestUserService struct {
	userServiceApiUrl string
}

type userResponseDto struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewRestUserService(
	userServiceApiUrl string,
) *RestUserService {
	return &RestUserService{
		userServiceApiUrl: userServiceApiUrl,
	}
}

func (r *RestUserService) GetUserById(userId int64) *_type.User {
	response, err := http.Get(r.userServiceApiUrl + "/api/v1/users/" + strconv.FormatInt(userId, 10))
	if err != nil {
		fmt.Printf("Error making the request: %v\n", err)
		return nil
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with status code: %d\n", response.StatusCode)
		return nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading the response body: %v\n", err)
		return nil
	}

	var user userResponseDto
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return nil
	}

	return &_type.User{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	}
}
