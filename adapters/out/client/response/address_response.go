package response

type AddressResponse struct {
	Cep         string `json:"cep"`
	Rua         string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"localidade"`
	Estado      string `json:"uf"`
}
