package repository

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(ctx context.Context, mongodb mongo.Collection, user dto.UserEntity) error {
	_, err := mongodb.InsertOne(ctx, user)
	return err
}
