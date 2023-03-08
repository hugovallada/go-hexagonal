package mock

import (
	"context"
	"errors"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity/model"
)

type DBMockWithoutError struct {
}

func (d *DBMockWithoutError) Execute(ctx context.Context, user model.UserModel) error {
	return nil
}

type DBMockWithError struct {
}

func (d *DBMockWithError) Execute(ctx context.Context, user model.UserModel) error {
	return errors.New("error while persisting")
}
