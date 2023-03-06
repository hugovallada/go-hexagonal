package mock

import (
	"context"
	"errors"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
)

type AddapterMockWithSuccessResponse struct{}

func (a *AddapterMockWithSuccessResponse) Execute(ctx context.Context, cep string) (dto.AddressFromCep, error) {
	return entity.Address{
		Cep:         "122",
		Rua:         "sda",
		Complemento: "sds",
		Bairro:      "sdsd",
		Cidade:      "de",
		Estado:      "sds",
	}, nil
}

type AddapterMockWithFailureResponse struct{}

func (a *AddapterMockWithFailureResponse) Execute(ctx context.Context, cep string) (dto.AddressFromCep, error) {
	return nil, errors.New("error while calling the external gateway")
}
