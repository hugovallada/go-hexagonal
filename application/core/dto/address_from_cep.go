package dto

type AddressFromCep interface {
	GetCep() string
	GetRua() string
	GetComplemento() string
	GetBairro() string
	GetCidade() string
	GetEstado() string
}
