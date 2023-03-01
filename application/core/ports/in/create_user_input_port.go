package in

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
)

type CreateUserInputPort interface {
	Execute(ctx context.Context, newUser dto.NewUser) (*entity.User, error)
}
