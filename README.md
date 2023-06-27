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

## Criação da tabela 

## Criação da tabela transactions no PostgreSQL

A tabela `transactions` representa o modelo de transação armazenado no banco de dados.

```sql
CREATE TABLE IF NOT EXISTS public.transactions
(
    id SERIAL PRIMARY KEY,
    cardholder VARCHAR(255),
    merchant VARCHAR(255),
    acquirer VARCHAR(255),
    brand VARCHAR(255),
    issuer VARCHAR(255)
);
```

E povoe o seu banco de dados, ou utilize o backup do diretório database-export

## Disclaimer

Antes de rodar a aplicação verifique se as configurações de banco de dados estão corretas no arquivo `dbConfig/configs.go`.

## Rodando o servidor do CRUD

```prompt 
PS C:\Users\User\go\src\github.com\lcslima45\transactions> cd transaction_main
PS C:\Users\User\go\src\github.com\lcslima45\transactions\transaction_main> go run main.go
Inicializing transaction server...
```

## Rodando o servidor de Autorização

```prompt 
PS C:\Users\User\go\src\github.com\lcslima45\transactions> cd authorization_main
PS C:\Users\User\go\src\github.com\lcslima45\transactions\authorization_main> go run main.go
Initializing authorization server...
```
## Testando uma criação de tabela autorizada 

Abra o Powershell e faça os seguintes passos para abrir uma requisição http de criação de uma nova tabela. Se a requisição atender aos critérios de autorização a resposta do servidor será positiva

```prompt
PS C:\Users\User\go\src\github.com\lcslima45> Invoke-WebRequest -Uri http://localhost:8080/transactions/add -Method POST -Body '{
>>     "cardholder": "Segio Lopes",
>>     "merchant": "Super Lagoa",
>>     "acquirer": "Caixa",
>>     "brand": "Visa",
>>     "issuer": "Bradesco"
>> }'


StatusCode        : 201
StatusDescription : Created
Content           : Transaction succesfully inserted!!!
RawContent        : HTTP/1.1 201 Created
                    Content-Length: 35
                    Content-Type: text/plain; charset=utf-8
                    Date: Tue, 27 Jun 2023 01:40:41 GMT

                    Transaction succesfully inserted!!!
Forms             : {}
Headers           : {[Content-Length, 35], [Content-Type, text/plain; charset=utf-8], [Date, Tue, 27 Jun 2023 01:40:41 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 35
```

## Caso de criação de tabela não autorizado

Caso a requisição possuam o mesmo Cardholder, Brand e Issuer de uma transação que já exista no banco de daods então a operação não será autorizada.

```prompt
PS C:\Users\User\go\src\github.com\lcslima45> Invoke-WebRequest -Uri http://localhost:8080/transactions/add -Method POST -Body '{
>>     "cardholder": "Segio Lopes",
>>     "merchant": "Super Lagoa",
>>     "acquirer": "Caixa",
>>     "brand": "Visa",
>>     "issuer": "Bradesco"
>> }'
Invoke-WebRequest : Transaction unauthorized
No linha:1 caractere:1
+ Invoke-WebRequest -Uri http://localhost:8080/transactions/add -Method ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebRequest) [Invoke-WebRequest], WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
```
## Caso onde uma tabela é excluída com sucesso 

Para deletar uma tabela basta executar uma requisição http no prompt de comando com o id
da tabela cuja exclusão é desejada.

```prompt 
PS C:\Users\User\go\src\github.com\lcslima45> Invoke-WebRequest -Uri http://localhost:8080/transactions/delete/36 -Method DELETE


StatusCode        : 200
StatusDescription : OK
Content           : Transaction succesfully deleted!!!
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 34
                    Content-Type: text/plain; charset=utf-8
                    Date: Tue, 27 Jun 2023 01:54:34 GMT

                    Transaction succesfully deleted!!!
Forms             : {}
Headers           : {[Content-Length, 34], [Content-Type, text/plain; charset=utf-8], [Date, Tue, 27 Jun 2023 01:54:34 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 34
```
## Caso onde a exclusão não é a autorizada

Uma tabela não tem a exclusão autorizada quando a Brand é igual a Visa, é o caso do id=37 da requisição bem sucedida do caso de criação de tabela que foi autorizado nesse Read.md nos passos anteriores, cujo cardholder = Segio Lopes e a brand = Visa. Veja como essa tabela não pode ser excluída 

```prompt 
PS C:\Users\User\go\src\github.com\lcslima45> Invoke-WebRequest -Uri http://localhost:8080/transactions/delete/37 -Method DELETE
Invoke-WebRequest : Deletion unauthorized
No linha:1 caractere:1
+ Invoke-WebRequest -Uri http://localhost:8080/transactions/delete/37 - ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebRequest) [Invoke-WebRequest], WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
```
## Caso onde o update é autorizada

Para fazer um update autorizado basta enviar um json no body da requisição http para o id da tabela cuja auteração é desejada. Ainda usando o exemplo da tabela de id=37

```prompt
PS C:\Users\User\go\src\github.com\lcslima45> Invoke-WebRequest -Uri http://localhost:8080/transactions/update/37 -Method PUT -Body '{
>>     "id": 37,
>>     "cardholder": "Segio Lopes",
>>     "merchant": "Extra",
>>     "acquirer": "Caixa",
>>     "brand": "Visa",
>>     "issuer": "Bradesco"
>> }'


StatusCode        : 200
StatusDescription : OK
Content           : Transaction succesfully updated!!!
RawContent        : HTTP/1.1 200 OK
                    Content-Length: 34
                    Content-Type: text/plain; charset=utf-8
                    Date: Tue, 27 Jun 2023 02:09:34 GMT

                    Transaction succesfully updated!!!
Forms             : {}
Headers           : {[Content-Length, 34], [Content-Type, text/plain; charset=utf-8], [Date, Tue, 27 Jun 2023 02:09:34 GMT]}
Images            : {}
InputFields       : {}
Links             : {}
ParsedHtml        : mshtml.HTMLDocumentClass
RawContentLength  : 34

```
## Caso onde o update não é autorizada

Caso você queira modificar o nome de uma tabela existente, a operação não será autorizada. Vamos tentar modificar o cardholder = Segio Lopes para Leonardo DiCaprio, a operação não será autorizada. 

```prompt 
PS C:\Users\User\go\src\github.com\lcslima45> Invoke-WebRequest -Uri http://localhost:8080/transactions/update/37 -Method PUT -Body '{
>>     "id": 37,
>>     "cardholder": "Leonardo DiCaprio",
>>     "merchant": "Extra",
>>     "acquirer": "Caixa",
>>     "brand": "Visa",
>>     "issuer": "Bradesco"
>> }'
Invoke-WebRequest : Update unauthorized
No linha:1 caractere:1
+ Invoke-WebRequest -Uri http://localhost:8080/transactions/update/37 - ...
+ ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
    + CategoryInfo          : InvalidOperation: (System.Net.HttpWebRequest:HttpWebRequest) [Invoke-WebRequest], WebException
    + FullyQualifiedErrorId : WebCmdletWebResponseException,Microsoft.PowerShell.Commands.InvokeWebRequestCommand
```

