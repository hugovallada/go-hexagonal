package models

type UserRequest struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	Birthday  string `json:"birthday"`
	Cellphone string `json:"cellphone"`
	Cep       string `json:"cep"`
}

func (u *UserRequest) GetName() string {
	return u.Name
}

func (u *UserRequest) GetLastName() string {
	return u.LastName
}

func (u *UserRequest) GetBirthDay() string {
	return u.Birthday
}

func (u *UserRequest) GetCellphone() string {
	return u.Cellphone
}

func (u *UserRequest) GetCep() string {
	return u.Cep
}
