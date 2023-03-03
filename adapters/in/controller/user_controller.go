package controller

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/in/controller/models"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/ports/in"
	"github.com/hugovallada/golang-hexagonal-architecture/configuration/logger"
	"go.uber.org/zap"
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
	logger.Info("Starting the operation to create user", zap.Time("time", time.Now()))
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
	logger.Info("Successfully created the new user", zap.String("name", userRequest.Name), zap.Time("time", time.Now()))
	ctx.JSON(201, user)
}
