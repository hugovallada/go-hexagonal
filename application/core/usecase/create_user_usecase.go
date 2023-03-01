package usecase

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto/converter"
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

func (u *CreateUserUseCase) Execute(ctx context.Context, newUser dto.NewUser) (*entity.User, error) {
	user, err := converter.FromNewUserDtoToUser(newUser)
	if err != nil {
		return nil, err
	}
	addressFromCep, err := u.getAddressByCepOutputPort.Execute(ctx, newUser.GetCep())
	if err != nil {
		return nil, err
	}
	address := converter.FromAddressFromCepToAddress(addressFromCep)
	user.AddAddress(address)
	if err := u.persistUserInDatabaseOutputPort.Execute(ctx, *user); err != nil {
		return nil, err
	}
	return user, nil
}
