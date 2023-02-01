
# Teste Golang

Aplicação de Teste que efutua cadastro de usuários e realiza transações entre eles.


## Rodando localmente

Clone o projeto

```bash
  git clone https://github.com/fransqueiroz/go_teste
```

Entre no diretório do projeto

```bash
  cd go_teste
```

Crie o arquivo .env com a URL do MOCK para autenticação

```json
MOCK_URL = ""
```

Inicie a aplicação com

```bash
  docker-compose up --build
```


## Exemplos de JSON para criações

Usuários
```json
POST: localhost:5000/user
{
    "name": "Nome do Usuário",
    "cpf": "xxxxxxxxxxx",
    "email": "email@email.com",
    "password": "123",
    "user_type": "F"
}
```

Adição de Valor na Carteira do Usuário

```json
PUT: localhost:5000/wallet/{user_id}
{
    "value": 1160.52
}
```

Transação
```json
POST: localhost:5000/transaction
{
    "value" : 20.00,
    "payer" : 1,
    "payee" : 2
}
```

