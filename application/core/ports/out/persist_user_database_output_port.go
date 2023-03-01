package out

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
)

type PersistUserInDatabaseOutputPort interface {
	Execute(ctx context.Context, user entity.User) error
}
