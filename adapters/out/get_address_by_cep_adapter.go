package out

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hugovallada/golang-hexagonal-architecture/adapters/out/client/converter"
	clientResp "github.com/hugovallada/golang-hexagonal-architecture/adapters/out/client/response"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/entity"
)

const (
	METHOD = "GET"
)

type GetAddressByCepAdapter struct{}

func (a *GetAddressByCepAdapter) Execute(ctx context.Context, cep string) (*entity.Address, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep), nil)
	if err != nil {
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	var addressResponse clientResp.AddressResponse
	err = json.NewDecoder(response.Body).Decode(&addressResponse)
	if err != nil {
		return nil, err
	}
	address := converter.FromAddressResponseToAddress(addressResponse)
	return &address, nil
}
