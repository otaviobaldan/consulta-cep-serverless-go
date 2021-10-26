package constant

const (
	ProviderBrasilApi = "brasil-api"
	ProviderViaCep    = "via-cep"
	ProviderBuscaCep  = "busca-cep"

	ProviderDefault = ProviderBrasilApi
)

var ProvidersAll = []string{ProviderViaCep, ProviderBrasilApi, ProviderBuscaCep}
