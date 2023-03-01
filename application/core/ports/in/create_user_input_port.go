package in

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
)

type CreateUserInputPort interface {
	Execute(ctx context.Context, user entity.User, cep string) (*entity.User, error)
}
