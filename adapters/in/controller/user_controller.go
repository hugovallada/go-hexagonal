package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/in/controller/models"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/ports/in"
	"github.com/hugovallada/golang-hexagonal-architecture/configuration/logger"
	"go.uber.org/zap"
)

const (
	FILENAME = "user_controller"
)

type UserInputPort interface {
	CreateUser(ctx *gin.Context)
}

type UserController struct {
	createUserInputPort in.CreateUserInputPort
}

func NewUserController(createUserInputPort in.CreateUserInputPort) *UserController {
	return &UserController{
		createUserInputPort: createUserInputPort,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	logger.Info("Starting the operation to create user", zap.Any("info", logger.NewLog[string, any, error](logger.NewJourney("createUser", ""), "CreateUser", FILENAME, "Request to create new user", ctx.Err() != nil, "Start", nil, nil)))
	var userRequest *models.UserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.AbortWithStatusJSON(400, "error while converting data to json")
		return
	}
	user, err := c.createUserInputPort.Execute(ctx, userRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(500, "error while creating the new user")
		return
	}
	logger.Info("User created...", zap.Any("info", logger.NewLog[*models.UserRequest, dto.UserEntity, error](logger.NewJourney("createUser", ""), "CreateUser", FILENAME, "Request to create new user", false, userRequest, user, nil)))
	ctx.JSON(201, user)
}
