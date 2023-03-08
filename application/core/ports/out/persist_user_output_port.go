package out

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity/model"
)

type PersistUserOutputPort interface {
	Execute(ctx context.Context, user model.UserModel) error
}
