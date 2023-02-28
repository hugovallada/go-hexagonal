package out

import "github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"

type GetAddressByCepOutputPort interface {
	Execute(cep string) (entity.Address, error)
}
