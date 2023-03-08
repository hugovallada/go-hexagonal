package model

import "time"

type UserModel interface {
	GetName() string
	GetLastName() string
	GetCellphone() string
	GetAddress() AddressModel
	GetAge() int8
	GetBirthDay() time.Time

	SetName(string)
	SetLastName(string)
	SetCellphone(string)
	SetAddress(AddressModel)
	SetBirthday(time.Time)
}
