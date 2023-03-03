package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/in/controller"
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/out"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/usecase"
	"github.com/hugovallada/golang-hexagonal-architecture/configuration/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:example@localhost:27017"))
	if err != nil {
		log.Fatal("error trying to connect to database")
		return
	}
	collection := client.Database("users").Collection("users")
	persistAdapter := out.PersistUserAdapter{
		Mongodb: *collection,
	}
	router := gin.Default()
	adp := out.GetAddressByCepAdapter{}
	uc := usecase.NewCreateUserUseCase(&adp, &persistAdapter)
	c := controller.NewUserController(uc)
	routes.InitRoutes(&router.RouterGroup, c)
	if err := router.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}
