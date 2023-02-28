package in

import "github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"

type CreateUserInputPort interface {
	Execute(user entity.User, cep string) (*entity.User, error)
}
