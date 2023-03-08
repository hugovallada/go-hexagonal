package entity

import (
	"time"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity/model"
)

type User struct {
	name      string
	lastName  string
	address   Address
	cellphone string
	birthday  time.Time
}

func NewUser(name, lastName, cellphone string, address Address, birthday time.Time) *User {
	return &User{
		name:      name,
		lastName:  lastName,
		cellphone: cellphone,
		address:   address,
		birthday:  birthday,
	}
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetLastName() string {
	return u.lastName
}

func (u User) GetCellphone() string {
	return u.cellphone
}

func (u User) GetAddress() model.AddressModel {
	return &u.address
}

func (u User) GetAge() int8 {
	ageInHours := time.Since(u.birthday).Hours()
	return int8(ageInHours / (24 * 365))
}

func (u User) GetBirthDay() time.Time {
	return u.birthday
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetLastName(lastName string) {
	u.lastName = lastName
}

func (u *User) SetBirthday(birthday time.Time) {
	u.birthday = birthday
}

func (u *User) SetCellphone(cellphone string) {
	u.cellphone = cellphone
}

func (u *User) SetAddress(address model.AddressModel) {
	u.address = *NewAddress(address.GetCep(), address.GetRua(), address.GetComplemento(), address.GetBairro(), address.GetCidade(), address.GetEstado())
}
