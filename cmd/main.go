package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/in/controller"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/usecase"
	"github.com/hugovallada/golang-hexagonal-architecture/config/routes"
)

func main() {
	router := gin.Default()
	uc := usecase.NewCreateUserUseCase(nil, nil)
	c := controller.NewUserController(uc)
	routes.InitRoutes(&router.RouterGroup, c)
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
