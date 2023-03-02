package dto

type NewAddress interface {
	SetCep(string)
	SetRua(string)
	SetComplemento(string)
	SetBairro(string)
	SetCidade(string)
	SetEstado(string)
	GetCep() string
	GetRua() string
	GetComplemento() string
	GetBairro() string
	GetCidade() string
	GetEstado() string
}
