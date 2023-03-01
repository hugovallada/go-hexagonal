package out

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
)

type GetAddressByCepOutputPort interface {
	Execute(ctx context.Context, cep string) (dto.AddressFromCep, error)
}
