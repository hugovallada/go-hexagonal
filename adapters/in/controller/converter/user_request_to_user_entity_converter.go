package converter

import (
	"github.com/hugovallada/golang-hexagonal-architecture/adapters/in/controller/models"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
)

func FromUserRequestToUserEntity(userRequest models.UserRequest) (*entity.User, error) {
	date, err := StringToDateConverter(userRequest.Birthday)
	if err != nil {
		return nil, err
	}
	return entity.NewUser(userRequest.Name, userRequest.LastName, userRequest.Cellphone, *&entity.Address{}, date), nil
}
