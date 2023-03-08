package in

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity/model"
)

type CreateUserInputPort interface {
	Execute(ctx context.Context, user dto.UserData) (model.UserModel, error)
}
