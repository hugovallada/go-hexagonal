package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/in/controller"
)

func InitRoutes(rg *gin.RouterGroup, controller controller.UserInputPort) {
	rg.POST("/users", controller.CreateUser)
}
