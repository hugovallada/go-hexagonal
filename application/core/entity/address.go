package entity

type Address struct {
	Cep         string
	Rua         string
	Complemento string
	Bairro      string
	Cidade      string
	Estado      string
}

func NewAddress(cep, rua, complemento, bairro, cidade, estado string) *Address {
	return &Address{
		Cep:         cep,
		Rua:         rua,
		Complemento: complemento,
		Bairro:      bairro,
		Cidade:      cidade,
		Estado:      estado,
	}
}
