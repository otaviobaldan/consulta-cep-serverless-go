package utils

import (
	"errors"
	"fmt"
	"github.com/otaviobaldan/consulta-cep-serverless-go/src/api/constant"
	"regexp"
	"strings"
)

func RemoveSpecialCharacter(cep string) string {
	return strings.Join(regexp.MustCompile(`\d`).FindAllString(cep, -1), "")
}

func ValidateCepLength(cep string) string {
	return PadLeft(cep, "0", constant.CepLength)
}

func PadLeft(str, pad string, length int) string {
	if len(str) == length {
		return str
	}
	for {
		if len(str)+1 > length {
			return str[0:length]
		} else {
			str = pad + str
		}
	}
}

func ValidateProvider(provider string) (string, error) {
	if provider == "" {
		return constant.ProviderDefault, nil
	}
	provider = strings.ToLower(provider)
	for i := range constant.ProvidersAll {
		if constant.ProvidersAll[i] == provider {
			return provider, nil
		}
	}

	return "", errors.New(fmt.Sprintf("The provider %s isn't compatible with that API.", provider))
}
