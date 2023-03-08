package model

type AddressModel interface {
	GetCep() string
	GetRua() string
	GetComplemento() string
	GetBairro() string
	GetCidade() string
	GetEstado() string

	SetCep(string)
	SetRua(string)
	SetComplemento(string)
	SetBairro(string)
	SetCidade(string)
	SetEstado(string)
}
