package out

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/adapters/out/repository"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity/model"
	"github.com/hugovallada/golang-hexagonal-architecture/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

const (
	FILENAME = "persist_user_adapter"
)

type PersistUserAdapter struct {
	Mongodb mongo.Collection
}

func (a *PersistUserAdapter) Execute(ctx context.Context, user model.UserModel) error {
	logger.Info("Persist User", zap.Any("execution", logger.NewLog[model.UserModel, any, error](logger.NewJourney("createUser", ""), "PersistUserAdapter.Execute()", FILENAME, "creating new user in the database", false, user, nil, nil)))
	err := repository.CreateUser(ctx, a.Mongodb, user)
	logger.Info("Persist User", zap.Any("execution", logger.NewLog(logger.NewJourney("createUser", ""), "PersistUserAdapter.Execute()", FILENAME, "creating new user in the database", err != nil, user, "", err)))
	return err
}
