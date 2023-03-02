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

func (ar Address) GetCep() string {
	return ar.Cep
}

func (ar Address) GetRua() string {
	return ar.Rua
}

func (ar Address) GetComplemento() string {
	return ar.Complemento
}

func (ar Address) GetBairro() string {
	return ar.Bairro
}

func (ar Address) GetCidade() string {
	return ar.Cidade
}
func (ar Address) GetEstado() string {
	return ar.Estado
}

func (ar *Address) SetCep(cep string) {
	ar.Cep = cep
}

func (ar *Address) SetRua(rua string) {
	ar.Rua = rua
}

func (ar *Address) SetComplemento(complemento string) {
	ar.Complemento = complemento
}

func (ar *Address) SetBairro(bairro string) {
	ar.Bairro = bairro
}

func (ar *Address) SetCidade(cidade string) {
	ar.Cidade = cidade
}
func (ar *Address) SetEstado(estado string) {
	ar.Estado = estado
}
