# transactions

## Estrutura de Pastas do Projeto
transactions/
├── dbConfig/
│   └── dbConfig.go
├── authorization/
│   ├── authorization_model.go
│   ├── authorization_handler.go
│   ├── authorization_repository.go
│   └── authorization_service.go
├── transaction/
│   ├── transaction_handler.go
│   ├── transaction_repository.go
│   ├── transaction_service.go
│   └── transaction_message.go
├── transaction_main/
│   └── main.go
├── authorization_main/
│   └── main.go
└── README.md



Aqui está a estrutura de pastas atualizada do projeto:

- `dbConfig/`: Contém o arquivo `dbConfig.go` que configura o banco de dados.
- `authorization/`: Contém os arquivos relacionados à autorização:
  - `authorization_model.go`: Define as estruturas de dados relacionadas ao modelo de autorização.
  - `authorization_handler.go`: Implementa os manipuladores HTTP para a autorização.
  - `authorization_repository.go`: Implementa a camada de acesso a dados para a autorização.
  - `authorization_service.go`: Implementa a lógica de negócios para a autorização.
- `transaction/`: Contém os arquivos relacionados às transações:
  - `transaction_handler.go`: Implementa os manipuladores HTTP para as transações.
  - `transaction_repository.go`: Implementa a camada de acesso a dados para as transações.
  - `transaction_service.go`: Implementa a lógica de negócios para as transações.
  - `transaction_message.go`: Define as funções de requisições HTTP que vão fazer chamadas ao microserviço de Autorização e retornam as autorizações para as funções de handler do transaction_handler.
- `transaction_main/`: Contém o arquivo `main.go` que inicia o servidor para as transações.
- `authorization_main/`: Contém o arquivo `main.go` que inicia o servidor para a autorização.

Certifique-se de substituir o conteúdo atual do README.md pelo esquema de pastas acima.
