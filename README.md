# Consulta CEP Serverless

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://img.shields.io/github/go-mod/go-version/otaviobaldan/consulta-cep-serverless-go)
[![GitHub license](https://img.shields.io/github/license/Naereen/StrapDown.js.svg)](https://github.com/Naereen/StrapDown.js/blob/master/LICENSE)

Consulta CEP foi desenvolvido com o objetivo de facilitar a vida do desenvolvedor que precisa de um serviço de consulta
de CEP, mas não quer dispor de tempo para desenvolver e nem de uma infra robusta para hospedar os recursos.

### Requisitos

- Golang 1.14 ou superior
- Serverless - [Guia de instalação](https://www.serverless.com/framework/docs/getting-started)
- AWS Keys configuradas - [doc](https://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/setup-credentials.html)

### Deploy / Instalação

```
    sh ./deploy.sh
```

**What!? É só isso?**

<img src="https://static.imgs.app/content/assetz/uploads/2017/06/meme-do.jpg" alt="what-meme" width="200"/>

Exatamente. O [Serverless](https://serverless.com/) será responsável por criar toda a infraestrutura, API Gateway e
fazer o deploy da função Lambda na AWS.

