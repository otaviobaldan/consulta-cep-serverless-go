package main

import (
    "context"
    "encoding/json"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/otaviobaldan/consulta-cep-serverless-go/src/api/adapter"
    "github.com/otaviobaldan/consulta-cep-serverless-go/src/api/constant"
    "github.com/otaviobaldan/consulta-cep-serverless-go/src/api/model"
    "github.com/otaviobaldan/consulta-cep-serverless-go/src/api/utils"
    "net/http"
    "strings"
)

func handleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    var err error
    var response *model.CepResponse

    cep := req.QueryStringParameters["cep"]
    cepSanitized := utils.RemoveSpecialCharacter(cep)
    cepSanitized = utils.ValidateCepLength(cepSanitized)

    provider := req.QueryStringParameters["provider"]
    provider, err = utils.ValidateProvider(provider)
    if err != nil {
        return buildAPIGatewayResponse(string(marshallError(err)), http.StatusBadRequest), nil
    }


    switch provider {
    case constant.ProviderBuscaCep:
        response, err = adapter.GetCepBuscaCep(cepSanitized)
    case constant.ProviderViaCep:
        response, err = adapter.GetCepViaCep(cepSanitized)
    case constant.ProviderBrasilApi:
        response, err = adapter.GetCepBrasilApi(cepSanitized)
    }
    if err != nil {
        if strings.Contains(err.Error(), "not found") {
           return buildAPIGatewayResponse(string(marshallError(err)), http.StatusNotFound), nil
        }
        return buildAPIGatewayResponse(string(marshallError(err)), http.StatusFailedDependency), nil
    }


    return buildAPIGatewayResponse(string(marshallResponse(response)), http.StatusOK), nil
}

func marshallResponse(response *model.CepResponse) []byte {
    bytes, _ := json.Marshal(response)
    return bytes
}

func marshallError(err error) []byte {
    type errStruct struct {
        Error string `json:"error"`
    }
    bytes, _ := json.Marshal(errStruct{Error: err.Error()})
    return bytes
}

func buildAPIGatewayResponse(body string, statusCode int) events.APIGatewayProxyResponse {
    return events.APIGatewayProxyResponse{
        Body: body,
        Headers: map[string]string{
            "Content-type": "application/json",
        },
        StatusCode: statusCode,
    }
}

func main() {
    lambda.Start(handleRequest)
}