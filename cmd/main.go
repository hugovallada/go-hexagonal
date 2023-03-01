package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/in/controller"
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/out"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/usecase"
	"github.com/hugovallada/golang-hexagonal-architecture/config/routes"
)

func main() {
	router := gin.Default()
	adp := out.GetAddressByCepAdapter{}
	uc := usecase.NewCreateUserUseCase(&adp, nil)
	c := controller.NewUserController(uc)
	routes.InitRoutes(&router.RouterGroup, c)
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
