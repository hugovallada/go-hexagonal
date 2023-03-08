package models

import (
	"fmt"

	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity/model"
)

type UserResponse struct {
	Name            string `json:"nome_completo"`
	Age             int8   `json:"idade_atual"`
	Cellphone       string `json:"numero_de_celular"`
	AddressResponse `json:"endereço_de_contato"`
}

func NewUserResponse(user model.UserModel) UserResponse {
	return UserResponse{
		Name:            fmt.Sprintf("%s %s", user.GetName(), user.GetLastName()),
		Age:             user.GetAge(),
		Cellphone:       user.GetCellphone(),
		AddressResponse: NewAddress(user.GetAddress()),
	}
}

type AddressResponse struct {
	Cep         string `json:"cep"`
	Rua         string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Cidade      string `json:"localidade"`
	Estado      string `json:"uf"`
}

func NewAddress(address model.AddressModel) AddressResponse {
	return AddressResponse{
		Cep:         address.GetCep(),
		Rua:         address.GetRua(),
		Complemento: address.GetComplemento(),
		Bairro:      address.GetBairro(),
		Cidade:      address.GetCidade(),
		Estado:      address.GetEstado(),
	}
}
