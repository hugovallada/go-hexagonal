package mock

import (
	"context"
	"errors"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
)

type DBMockWithoutError struct {
}

func (d *DBMockWithoutError) Execute(ctx context.Context, user dto.UserEntity) error {
	return nil
}

type DBMockWithError struct {
}

func (d *DBMockWithError) Execute(ctx context.Context, user dto.UserEntity) error {
	return errors.New("error while persisting")
}
