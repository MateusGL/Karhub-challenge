# Karhub-challenge

## Pré-requisitos

- Go 1.20 ou superior instalado
- SQLite3 instalado (para uso local)
- Conta no [Spotify Developer](https://developer.spotify.com/) para obter um token de acesso

## Configuração do ambiente

1. **Clone o repositório:**
   ```sh
   git clone https://github.com/seu-usuario/Karhub-challenge.git
   cd Karhub-challenge
   ```

2. **Crie um arquivo .env na raiz do projeto com as variáveis necessárias:**
    ```
    =SPOTIFY_ACCESS_TOKENseu_token_aqui
    PORT=8080
    ```

- O SPOTIFY_ACCESS_TOKEN pode ser obtido em Spotify Developer Console.
- O token deve ter permissão para acessar as APIs públicas do Spotify.


3. **Instale as dependências do projeto**
    ```
    go mod tidy
    ```

4. **(Opcional) Instale o SQLite3 caso não tenha:**

Rodando o projeto

Execute o comando abaixo na raiz do projeto

    ```sh
    go run [main.go](http://_vscodecontentref_/0)
    ```

# Comandos para testar as rotas

1. **Listar todos os estilos de cerveja (GET /beers)**
    ```
    curl -X GET http://localhost:8080/beers
    ```

2. **Criar um novo estilo de cerveja (POST /beers)**
    ```
    curl -X POST http://localhost:8080/beers \
    -H "Content-Type: application/json" \
    -d '{
      "name": "IPA",
      "minTemp": -7,
      "maxTemp": 10
    }'
    ```

3. **Atualizar um estilo de cerveja pelo ID (PUT /beers/{id})**
    ```
    curl -X PUT http://localhost:8080/beers/1 \
    -H "Content-Type: application/json" \
    -d '{
      "name": "Imperial IPA",
      "minTemp": -8,
      "maxTemp": 11
    }'
    ```

4. **Deletar um estilo de cerveja pelo ID (DELETE /beers/{id})**
    ```
    curl -X DELETE http://localhost:8080/beers/1
    ```

5. **Recomendação de estilo pela temperatura (POST /recommendation)**
    ```
    curl -X POST http://localhost:8080/recommendation \
    -H "Content-Type: application/json" \
    -d '{
      "temperature": -6
    }'
    ```