## Como usar

- Primeiro é necessário que tenha o golang(https://golang.org/doc/install) e o docker (https://docs.docker.com/get-docker/) instalados.

- Para rodar a aplicação entre na pasta `pismo` e excute o seguinte comando `docker-compose up -d`.

- Após isso aplique a migração inicial `migrate -path db/migration -database "postgresql://pismo:password@localhost:5432/pismo_api?sslmode=disable" -verbose up`

- Após isso a aplicação estará rodando na url: `localhost:8080`

## Endpoints

- A aplicação possui 3 endpoints
    - POST `localhost:8080/accounts/create`
        - Request Body:
            ```shell
                { 
                    "document_number": "12345678900" 
                }

            ```

        - Response exemplo:

            ```shell
                {
                    "ID": 1,
                    "DocumentNumber": "12345678900",
                    "Balance": 0,
                    "CreatedAt": "2020-11-02T20:08:45.381641Z"
                }

            ```

    - GET `localhost:8080/accounts/:id`
        - Request exemplo:
            ```shell
                localhost:8080/accounts/1:
            ```

        - Response exemplo:

            ```shell
                {
                    "ID": 1,
                    "DocumentNumber": "12345678900",
                    "Balance": 123.45,
                    "CreatedAt": "2020-11-02T20:08:45.381641Z"
                }

            ```

    - POST `localhost:8080/transactions`
        - Request exemplo:
            ```shell
                { 
                    "account_id": 1, 
                    "operation_type_id": 4, 
                    "amount": 123.45 
                }
            ```

        - Response exemplo:

            ```shell
                {
                    "AccountID": 1,
                    "Amount": 123.45,
                    "TransactionType": "payment",
                    "EventDate": "2020-11-02T20:11:26.964433Z"
                }

            ```