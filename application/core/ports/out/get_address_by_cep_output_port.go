package out

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
)

type GetAddressByCepOutputPort interface {
	Execute(ctx context.Context, cep string) (*entity.Address, error)
}
