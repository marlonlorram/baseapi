# baseapi

[![License](https://img.shields.io/github/license/marlonlorram/baseapi)](https://github.com/marlonlorram/baseapi/blob/main/LICENSE)
[![Last Commit](https://img.shields.io/github/last-commit/marlonlorram/baseapi)](https://github.com/marlonlorram/baseapi/commits/main)
[![go.mod Version](https://img.shields.io/github/go-mod/go-version/marlonlorram/baseapi)](https://github.com/marlonlorram/baseapi/blob/main/go.mod)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmarlonlorram%2Fbaseapi.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmarlonlorram%2Fbaseapi?ref=badge_shield)

![Alt](https://repobeats.axiom.co/api/embed/facf73357a293e9ee9e7404d6f96883999b6a07e.svg "Repobeats analytics image")

Este projeto é uma API RESTful construída com `Go`. Utilizamos o framework `Gin` para o roteamento HTTP, utilizando as bibliotecas `go.uber/fx` para injeção de dependência, `go.uber/zap` para logging e `MongoDB` como banco de dados.

## Tecnologias Utilizadas

- Go
- Gin
- go.uber/fx
- go.uber/zap
- MongoDB

## Instalação e Configuração

### Pré-requisitos
- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [go.uber/fx](https://github.com/uber-go/fx)
- [go.uber/zap](https://github.com/uber-go/zap)
- [MongoDB](https://www.mongodb.com/)

### Passos

1. **Clone o repositório**
    ```
    git clone https://github.com/marlonlorram/baseapi.git
    ```

2. **Navegue até o diretório do projeto**
    ```
    cd baseapi
    ```

3. **Instale as dependências**
    ```
    go mod download
    ```

4. **Execute o projeto**
    ```
    task build # Para compilar a aplicação
    task prod  # Para rodar em ambiente de produção
    task local # Para rodar em ambiente de desenvolvimento
    ```

## Contribuição

Seu interesse em contribuir é apreciado!

Para contribuir, por favor, crie um fork do projeto, faça suas alterações e abra um Pull Request.

## Licença

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fmarlonlorram%2Fbaseapi.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fmarlonlorram%2Fbaseapi?ref=badge_large)