package out

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	clientResp "github.com/hugovallada/golang-hexagonal-architecture/adapters/out/client/response"
	"github.com/hugovallada/golang-hexagonal-architecture/application/core/dto"
	"github.com/hugovallada/golang-hexagonal-architecture/configuration/logger"
	"go.uber.org/zap"
)

const (
	METHOD = "GET"
)

type GetAddressByCepAdapter struct{}

func (a *GetAddressByCepAdapter) Execute(ctx context.Context, cep string) (dto.AddressFromCep, error) {
	logger.Info("Calling the external gateway", zap.Any("execution", logger.NewLog[string, any, error](logger.NewJourney("createUser", ""), "GetAddressByCepAdapter.Execute()", "get_address_by_cep_adapter", "start of method", false, cep, nil, nil)))
	request, err := http.NewRequestWithContext(ctx, "GET", fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep), nil)
	if err != nil {
		logger.Error("error while parsing the request", err, zap.Any("execution", logger.NewLog(logger.NewJourney("createUser", ""), "GetAddressByCepAdapter.Execute()", "get_address_by_cep_adapter", "mount request", true, cep, "", err)))
		return nil, err
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		logger.Error("error while calling the api", err, zap.Any("execution", logger.NewLog(logger.NewJourney("createUser", ""), "GetAddressByCepAdapter.Execute()", "get_address_by_cep_adapter", "calling api", true, cep, "", err)))
		return nil, err
	}
	defer response.Body.Close()
	var address clientResp.AddressResponse
	err = json.NewDecoder(response.Body).Decode(&address)
	if err != nil {
		logger.Error("error while binding the json", err, zap.Any("execution", logger.NewLog(logger.NewJourney("createUser", ""), "GetAddressByCepAdapter.Execute()", "get_address_by_cep_adapter", "binding json", true, cep, "", err)))
		return nil, err
	}
	logger.Info("error while calling the api", zap.Any("execution", logger.NewLog(logger.NewJourney("createUser", ""), "GetAddressByCepAdapter.Execute()", "get_address_by_cep_adapter", "return", false, cep, address, err)))
	return address, nil
}
