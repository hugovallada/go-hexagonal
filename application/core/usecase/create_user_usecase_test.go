package usecase

import (
	"context"
	"testing"

	"github.com/hugovallada/golang-hexagonal-architecture/application/tests/mock"
	"github.com/stretchr/testify/assert"
)

func Test_UseCaseShouldReturnWithoutAnyError(t *testing.T) {
	adapter := mock.AddapterMockWithSuccessResponse{}
	db := mock.DBMockWithoutError{}
	uc := NewCreateUserUseCase(&adapter, &db)
	user := mock.UserEntityMock{
		Name:      "Hugo",
		LastName:  "Lopes",
		Birthday:  "16/12/1995",
		Cellphone: "10292992",
		Cep:       "14010090",
	}
	_, err := uc.Execute(context.Background(), user)
	assert.Nil(t, err)
}
