package converter

import (
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
)

func FromAddressFromCepToAddress(address dto.AddressFromCep) entity.Address {
	return *entity.NewAddress(address.GetCep(), address.GetRua(), address.GetComplemento(), address.GetBairro(), address.GetCidade(), address.GetEstado())
}
