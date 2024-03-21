# Projeto

## Estrutura de pastas

``` 
.
├── cmd
│   └── app
│       └── main.go
├── internal
│   ├── user
│   │   ├── delivery
│   │   │   ├── http
│   │   │   │   ├── handler.go
│   │   │   │   └── router.go
│   │   │   └── grpc
│   │   ├── repository
│   │   │   ├── psql
│   │   │   │   └── user.go
│   │   │   └── mongo
│   │   │       └── user.go
│   │   └── usecase
│   │       └── user.go
│   └── middleware
│       └── middleware.go
├── pkg
│   └── utils
│       └── utils.go
└── api
    └── proto
        └── user
            └── user.proto
```

Explicação:

- cmd: Contém o ponto de entrada da aplicação (main.go).
- internal: Contém o código da aplicação que é específico para este projeto.
- user: Este é um módulo do projeto, cada módulo representa um domínio na arquitetura limpa.
- delivery: Contém o código que lida com a entrega do HTTP ou gRPC, como manipuladores e roteadores.
- repository: Contém o código que lida com o armazenamento de dados, como consultas SQL ou comandos MongoDB.
- usecase: Contém a lógica de negócios da aplicação.
- middleware: Contém o código do middleware da aplicação.
- pkg: Contém o código que pode ser usado por aplicativos externos.
- api: Contém arquivos de definição de API, como arquivos protobuf.

Por favor, note que esta é apenas uma sugestão de estrutura de diretório. A estrutura de diretório pode variar dependendo das necessidades do projeto.

