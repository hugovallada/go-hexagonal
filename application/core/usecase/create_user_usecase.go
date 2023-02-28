package usecase

import (
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/ports/out"
)

type CreateUserUseCase struct {
	getAddressByCepOutputPort       out.GetAddressByCepOutputPort
	persistUserInDatabaseOutputPort out.PersistUserInDatabaseOutputPort
}

func NewCreateUserUseCase(getAddressByCepOutputPort out.GetAddressByCepOutputPort, persistUserInDatabaseOutputPort out.PersistUserInDatabaseOutputPort) *CreateUserUseCase {
	return &CreateUserUseCase{
		getAddressByCepOutputPort:       getAddressByCepOutputPort,
		persistUserInDatabaseOutputPort: persistUserInDatabaseOutputPort,
	}
}

func (u *CreateUserUseCase) Execute(user entity.User, cep string) (*entity.User, error) {
	address, err := u.getAddressByCepOutputPort.Execute(cep)
	if err != nil {
		return nil, err
	}
	user.AddAddress(address)
	if err := u.persistUserInDatabaseOutputPort.Execute(user); err != nil {
		return nil, err
	}
	return &user, nil
}
