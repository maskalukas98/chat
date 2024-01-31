package rest

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"user_service/src/application/command"
	custom_error "user_service/src/application/error"
	"user_service/src/domain/input_port"
	"user_service/src/infrastructure/rest/dto"
)

type UserController struct {
	createUserUseCase input_port.CreateUserPort
	getUserUseCase    input_port.GetUserPort
}

func NewUserController(
	createUserUseCase *input_port.CreateUserPort,
	getUserUseCase *input_port.GetUserPort,
) *UserController {
	return &UserController{
		createUserUseCase: *createUserUseCase,
		getUserUseCase:    *getUserUseCase,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @ID create-user
// @Produce json
// @Param request body dto.CreateUserRequestDto true "User creation request"
// @Success 201 {object} dto.CreateUserResponseDto
// @Failure 400 {object} error "Bad Request"
// @Failure 409 {object} error "Conflict - User already exists"
// @Failure 500 {object} error "Internal Server Error"
// @Router /api/v1/users [post]
func (r *UserController) CreateUser(c *gin.Context) {
	var request = command.CreateUserRequest{
		IpAddress: c.ClientIP(),
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := r.createUserUseCase.Create(request)

	if err != nil {
		var userExistsErr *custom_error.UserAlreadyExistsError
		if errors.As(err, &userExistsErr) {
			c.JSON(http.StatusConflict, gin.H{
				"error_message": fmt.Sprintf("User already exists with email: %s", request.Email),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Internal server error"))
		return
	}

	c.JSON(http.StatusCreated, dto.CreateUserResponseDto{
		UserId: response.UserId,
		Links: map[string]string{
			"self": c.Request.URL.String() + "/" + strconv.FormatInt(response.UserId, 10),
		},
	})
}

// GetUserById godoc
// @Summary Get user by ID
// @Description Get user information by ID
// @ID get-user-by-id
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.GetUserResponseDto
// @Failure 400 {object} error "Bad Request"
// @Failure 404 {object} error "Not Found"
// @Failure 500 {object} error "Internal Server Error"
// @Router /api/v1/users/{id} [get]
func (r *UserController) GetUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	request := command.GetUserByIdRequest{UserId: int64(userId)}
	response, err := r.getUserUseCase.GetById(request)

	if response == nil && err == nil {
		// todo: create custom error
		c.JSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("User with ID %d not found.", userId),
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("Internal server error"))
		return
	}

	c.JSON(http.StatusOK, dto.GetUserResponseDto{
		UserId: response.UserId,
		Name:   response.Name,
		Email:  response.Email,
	})
}

func (r *UserController) GetMethodsForUsers(c *gin.Context) {
	c.Header("Allow", "POST")
	c.Status(http.StatusOK)
}

func (r *UserController) GetMethodsForOneUser(c *gin.Context) {
	c.Header("Allow", "GET")
	c.Status(http.StatusOK)
}
