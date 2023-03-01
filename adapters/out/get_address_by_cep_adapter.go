package out

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	clientResp "github.com/hugovallada/golang-hexagonal-architecture/adapters/out/client/response"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
)

const (
	METHOD = "GET"
)

type GetAddressByCepAdapter struct{}

func (a *GetAddressByCepAdapter) Execute(ctx context.Context, cep string) (dto.AddressFromCep, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep), nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	var address clientResp.AddressResponse
	err = json.NewDecoder(response.Body).Decode(&address)
	if err != nil {
		return nil, err
	}
	return address, nil
}
