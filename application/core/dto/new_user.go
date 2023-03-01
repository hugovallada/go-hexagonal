package dto

type NewUser interface {
	GetName() string
	GetLastName() string
	GetBirthday() string
	GetCellphone() string
	GetCep() string
}
