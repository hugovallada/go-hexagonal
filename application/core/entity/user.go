package entity

import "time"

type User struct {
	Name     string
	LastName string
	Address
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
