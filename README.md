# URL Shortener

Esta é uma aplicação de encurtador de URLs escrita em Go. A aplicação pode ser executada tanto via linha de comando quanto via API.

## Pré-requisitos

- [Go](https://golang.org/doc/install) 1.19 ou superior
- [Docker](https://docs.docker.com/get-docker/)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Instalação

### Clonar o repositório

Primeiro, clone o repositório para o seu ambiente local:

```sh
git clone https://github.com/DevLeoo/url-shortener.git
cd url-shortener
```

### Instalar pacotes necessários

Instale os pacotes necessários usando o comando `go mod`:

```sh
go mod download
```

## Executar a aplicação

A diferenciação entre execução da CLI ou da API é feita pelas variveis de ambiente.

1. Crie um arquivo .env seguindo o modelo apresentado no arquivo.env.example
2. Para a variavel ENV, escolha LOCAL para executar a CLI ou qualquer outra opção para executar a API.

### Via linha de comando
```sh
go run main.go shorten --urls="https://example.com,https://another.com"
```

### Via API

Para executar a aplicação via API, use o comando:

```sh
go run main.go
```

A aplicação estará disponível em `http://localhost:8080`.

## Consumir os recursos

### Linha de comando

#### Encurtar URLs

Para encurtar URLs via linha de comando, use o comando:

```sh
go run main.go shorten --urls="https://example.com,https://another.com"
```

#### Redirecionar URLs

Para redirecionar URLs via linha de comando, use o comando:

```sh
go run main.go key --keys="shortKey1,shortKey2"
```

### API

#### Encurtar URLs

Para encurtar URLs via API, envie uma solicitação POST para `http://localhost:8080/shorten` com um corpo JSON contendo uma lista de URLs:

```sh
curl -X POST -H "Content-Type: application/json" -d '["https://example.com", "https://another.com"]' http://localhost:8080/shorten
```

#### Redirecionar URLs

Para redirecionar URLs via API, envie uma solicitação GET para `http://localhost:8080/{shortKey}`, onde `{shortKey}` é a chave curta gerada.

## Estrutura do Projeto

```
url-shortener/
├── app/
│   ├── api/
│   │   ├── controllers/
│   │   │   └── shortener.go
│   │   ├── responses/
│   │   └── router/
│   │       └── routes/
│   │           └── shortener.go
│   ├── cli/
│   │   └── cli.go
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── redis.go
│   └── services/
│       └── shortener.go
├── .env
├── .gitignore
├── Dockerfile
├── go.mod
├── go.sum
└── main.go
```


