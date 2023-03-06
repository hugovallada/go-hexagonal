package mock

type UserEntityMock struct {
	Name      string 
	LastName  string 
	Birthday  string 
	Cellphone string 
	Cep       string 
}

func (u UserEntityMock) GetName() string {
	return u.Name
}

func (u UserEntityMock) GetLastName() string {
	return u.LastName
}

func (u UserEntityMock) GetBirthday() string {
	return u.Birthday
}

func (u UserEntityMock) GetCellphone() string {
	return u.Cellphone
}

func (u UserEntityMock) GetCep() string {
	return u.Cep
}