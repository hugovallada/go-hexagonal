package converter

import (
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
)

func FromNewUserDtoToUser(dto dto.NewUser) (*entity.User, error) {
	birthday, err := StringToDateConverter(dto.GetBirthday())
	if err != nil {
		return nil, err
	}
	return entity.NewUser(dto.GetName(), dto.GetLastName(), dto.GetCellphone(), entity.Address{}, birthday), nil
}
