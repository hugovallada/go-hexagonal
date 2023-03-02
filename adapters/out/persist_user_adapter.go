package out

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/adapters/out/repository"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersistUserAdapter struct {
	Mongodb mongo.Collection
}

func (a *PersistUserAdapter) Execute(ctx context.Context, user dto.UserEntity) error {
	err := repository.CreateUser(ctx, a.Mongodb, user)
	return err
}
