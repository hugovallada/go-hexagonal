package usecase

import (
	"context"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto/converter"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/ports/out"
)

type CreateUserUseCase struct {
	getAddressByCepOutputPort out.GetAddressByCepOutputPort
	persistUserOutputPort     out.PersistUserOutputPort
}

func NewCreateUserUseCase(getAddressByCepOutputPort out.GetAddressByCepOutputPort, persistUserOutputPort out.PersistUserOutputPort) *CreateUserUseCase {
	return &CreateUserUseCase{
		getAddressByCepOutputPort: getAddressByCepOutputPort,
		persistUserOutputPort:     persistUserOutputPort,
	}
}

func (u *CreateUserUseCase) Execute(ctx context.Context, receivedUser dto.NewUser) (dto.UserEntity, error) {
	user, err := converter.FromNewUserDtoToUser(receivedUser)
	if err != nil {
		return nil, err
	}
	addressFromCep, err := u.getAddressByCepOutputPort.Execute(ctx, receivedUser.GetCep())
	if err != nil {
		return nil, err
	}
	address := converter.FromAddressFromCepToAddress(addressFromCep)
	user.AddAddress(address)
	if err := u.persistUserOutputPort.Execute(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}
