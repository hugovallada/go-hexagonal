package in

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
)

type CreateUserInputPort interface {
	Execute(ctx context.Context, newUser dto.NewUser) (dto.UserEntity, error)
}
