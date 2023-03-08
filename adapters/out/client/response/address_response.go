package response

type AddressResponse struct {
	Cep         string `json:"cep"`
	Rua         string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"localidade"`
	Estado      string `json:"uf"`
}

func (ar AddressResponse) GetCep() string {
	return ar.Cep
}

func (ar AddressResponse) GetRua() string {
	return ar.Rua
}

func (ar AddressResponse) GetComplemento() string {
	return ar.Complemento
}

func (ar AddressResponse) GetBairro() string {
	return ar.Bairro
}

func (ar AddressResponse) GetCidade() string {
	return ar.Cidade
}
func (ar AddressResponse) GetEstado() string {
	return ar.Estado
}

func (ar *AddressResponse) SetCep(cep string) {
	ar.Cep = cep
}

func (ar *AddressResponse) SetRua(rua string) {
	ar.Rua = rua
}

func (ar *AddressResponse) SetComplemento(complemento string) {
	ar.Complemento = complemento
}

func (ar *AddressResponse) SetBairro(bairro string) {
	ar.Bairro = bairro
}

func (ar *AddressResponse) SetCidade(cidade string) {
	ar.Cidade = cidade
}
func (ar *AddressResponse) SetEstado(estado string) {
	ar.Estado = estado
}
