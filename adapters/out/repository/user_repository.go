package repository

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(ctx context.Context, mongodb mongo.Collection, user model.UserModel) error {
	_, err := mongodb.InsertOne(ctx, user)
	return err
}
