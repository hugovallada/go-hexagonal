package dto

import "time"

type UserEntity interface {
	SetName(name string)
	SetLastName(lastName string)
	SetBirthday(birthday time.Time)
	SetCellphone(cellphone string)
	SetAddress(address NewAddress)
}
