package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/in/controller/converter"
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/in/controller/models"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/ports/in"
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
	var userRequest *models.UserRequest
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.AbortWithStatusJSON(400, "error while converting data to json")
		return
	}
	user, err := converter.FromUserRequestToUserEntity(*userRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(400, "birthday in invalid format")
		return
	}
	user, err = c.createUserInputPort.Execute(ctx, *user, userRequest.Cep)
	if err != nil {
		ctx.AbortWithStatusJSON(500, "error while creating the new user")
		return
	}
	ctx.JSON(201, *user)
}
