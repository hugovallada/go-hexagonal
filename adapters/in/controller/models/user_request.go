package models

type UserRequest struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Birthday  string `json:"birthday"`
	Cellphone string `json:"cellphone"`
	Cep       string `json:"cep"`
}
