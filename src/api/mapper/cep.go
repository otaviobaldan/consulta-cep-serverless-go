package mapper

import "github.com/otaviobaldan/consulta-cep-serverless-go/src/api/model"

func MapViaApiToCepResponse(viaApiResponse model.CepViaApi) *model.CepResponse {
	return &model.CepResponse{
		Cep:          viaApiResponse.Cep,
		State:        viaApiResponse.Uf,
		City:         viaApiResponse.Localidade,
		Neighborhood: viaApiResponse.Bairro,
		Street:       viaApiResponse.Logradouro,
	}
}

func MapBuscaCepToCepResponse(buscaCepResponse model.CepBuscaCep) *model.CepResponse {
	return &model.CepResponse{
		Cep:          buscaCepResponse.Code,
		State:        buscaCepResponse.State,
		City:         buscaCepResponse.City,
		Neighborhood: buscaCepResponse.District,
		Street:       buscaCepResponse.Address,
	}
}
