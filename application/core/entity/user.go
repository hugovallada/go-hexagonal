package entity

import (
	"time"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
)

type User struct {
	Name      string
	LastName  string
	Address   Address
	Cellphone string
	Birthday  time.Time
}

func NewUser(name, lastName, cellphone string, address Address, birthday time.Time) *User {
	return &User{
		Name:      name,
		LastName:  lastName,
		Cellphone: cellphone,
		Address:   address,
		Birthday:  birthday,
	}
}

func (u User) GetAge() int8 {
	ageInHours := time.Since(u.Birthday).Hours()
	return int8(ageInHours / (24 * 365))
}

func (u *User) AddAddress(address Address) {
	u.Address = address
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetLastName(lastName string) {
	u.LastName = lastName
}

func (u *User) SetBirthday(birthday time.Time) {
	u.Birthday = birthday
}

func (u *User) SetCellphone(cellphone string) {
	u.Cellphone = cellphone
}

func (u *User) SetAddress(address dto.NewAddress) {
	u.AddAddress(*NewAddress(address.GetCep(), address.GetRua(), address.GetComplemento(), address.GetBairro(), address.GetCidade(), address.GetEstado()))
}
