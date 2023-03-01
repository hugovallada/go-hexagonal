package converter

import (
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/out/client/response"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
)

func FromAddressResponseToAddress(address response.AddressResponse) entity.Address {
	return *entity.NewAddress(address.Cep, address.Rua, address.Complemento, address.Bairro, address.Cidade, address.Estado)
}
