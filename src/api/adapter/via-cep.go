package adapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/otaviobaldan/consulta-cep-serverless-go/src/api/mapper"
	"github.com/otaviobaldan/consulta-cep-serverless-go/src/api/model"
	"github.com/otaviobaldan/consulta-cep-serverless-go/src/api/utils"
	"io/ioutil"
	"net/http"
	"time"
)

const viaCepUrl = "https://viacep.com.br/ws/%s/json/"

func GetCepViaCep(cep string) (*model.CepResponse, error) {
	var rest http.Client
	var viaApiResponse model.CepViaApi
	rest.Timeout = 5 * time.Second

	url := fmt.Sprintf(viaCepUrl, cep)

	response, err := rest.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error executing request to ViaCep API. Error: %s", err.Error()))
	}
	statusCode := response.StatusCode

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	if !utils.IsSuccessStatusCode(statusCode) || utils.IsBadRequestStatusCode(statusCode) {
		return nil, errors.New(fmt.Sprintf("Invalid response status code of ViaCep API. CEP: %s / Status Code: %d ",
			cep,
			statusCode,
		))
	}

	if utils.IsNotFoundStatusCode(statusCode) {
		return nil, errors.New(fmt.Sprintf("The provided CEP was not found. CEP: %s / Status Code: %d", cep, statusCode))
	}

	_ = json.Unmarshal(bodyBytes, &viaApiResponse)

	return mapper.MapViaApiToCepResponse(viaApiResponse), nil
}
