package models

import (
	"time"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSchema struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	LastName  string             `bson:"last_name"`
	Birthday  time.Time          `bson:"birthday"`
	Cellphone string             `bson:"cellphone"`
	Addresss  dto.AddressFromCep `bson:"address"`
}

func (u *UserSchema) SetName(name string) {
	u.Name = name
}

func (u *UserSchema) SetLastName(lastName string) {
	u.LastName = lastName
}

func (u *UserSchema) SetBirthday(birthday time.Time) {
	u.Birthday = birthday
}

func (u *UserSchema) SetCellphone(cellphone string) {
	u.Cellphone = cellphone
}

func (u *UserSchema) SetAddress(address dto.AddressFromCep) {
	u.Addresss = address
}
