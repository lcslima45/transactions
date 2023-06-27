# transactions

## Estrutura de Pastas do Projeto
Descrição dos diretórios do projeto:

- `dbConfig/`: Contém as configurações do banco de dados.
- `authorization/`: Contém os arquivos relacionados à autorização:
  - `authorization_model.go`: Define as estruturas de dados relacionadas ao modelo de autorização.
  - `authorization_handler.go`: Implementa os manipuladores HTTP para a autorização.
  - `authorization_repository.go`: Implementa a camada de acesso a dados para a autorização.
  - `authorization_service.go`: Implementa a lógica de negócios para a autorização.
- `transaction/`: Contém os arquivos relacionados às transações:
  - `transaction_handler.go`: Implementa os manipuladores HTTP para as transações.
  - `transaction_repository.go`: Implementa a camada de acesso a dados para as transações.
  - `transaction_service.go`: Implementa a lógica de negócios para as transações.
  - `transaction_message.go`: Define as funções de requisições HTTP que vão fazer chamadas ao microserviço de Autorização e retornam as autorizações para as funções de handler do `transaction_handler`.
- `transaction_main/`: Contém o arquivo `main.go` que inicia o servidor para as transações.
- `authorization_main/`: Contém o arquivo `main.go` que inicia o servidor para a autorização.

## Dependências

No prompt (Windows) ou Terminal (Linux) utilize esse comando:

```shell
go get github.com/lib/pq
```

## Download do PostgreSQL

Você pode baixar o PostgreSQL em [postgresql.org/download](https://www.postgresql.org/download/).


