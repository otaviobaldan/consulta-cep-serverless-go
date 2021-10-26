package adapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/otaviobaldan/consulta-cep-serverless-go/src/api/model"
	"github.com/otaviobaldan/consulta-cep-serverless-go/src/api/utils"
	"io/ioutil"
	"net/http"
	"time"
)

const brasilApiUrl = "https://brasilapi.com.br/api/cep/v1/%s"

func GetCepBrasilApi(cep string) (*model.CepResponse, error) {
	var rest http.Client
	var brasilApiResponse *model.CepResponse
	rest.Timeout = 5 * time.Second

	url := fmt.Sprintf(brasilApiUrl, cep)

	response, err := rest.Get(url)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error executing request to BrasilApi. Error: %s", err.Error()))
	}
	statusCode := response.StatusCode

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	if !utils.IsSuccessStatusCode(statusCode) || utils.IsBadRequestStatusCode(statusCode) {
		return nil, errors.New(fmt.Sprintf("Invalid response status code of BrasilApi. CEP: %s / Status Code: %d",
			cep,
			statusCode,
		))
	}

	if utils.IsNotFoundStatusCode(statusCode) {
		return nil, errors.New(fmt.Sprintf("The provided CEP was not found. CEP: %s / Status Code: %d", cep, statusCode))
	}

	_ = json.Unmarshal(bodyBytes, &brasilApiResponse)

	return brasilApiResponse, nil
}
