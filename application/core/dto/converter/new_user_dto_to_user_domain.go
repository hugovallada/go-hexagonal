package converter

import (
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
)

func FromNewUserDtoToUser(user dto.UserData) (*entity.User, error) {
	birthday, err := StringToDateConverter(user.GetBirthDay())
	if err != nil {
		return nil, err
	}
	return entity.NewUser(user.GetName(), user.GetLastName(), user.GetCellphone(), entity.Address{}, birthday), nil
}
