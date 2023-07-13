# Desafio Codigo - *REST API* que faça conversão de moedas em Golang.

Para iniciar, rode o comando `docker-compose up` para subir o container do projeto.

Estarei deixando o arquivo .env com uma pré configuração setada.

Estou utilizando como framework o Echo, base de dados Mysql com o [GORM](https://gorm.io/).

## Rotas:

* GetExchange:

    Aqui você envia o valor que deseja fazer a conversão, a moeda de origem, moeda para conversão e a cotação.

    http://localhost:8000/exchange/{amount}/{from}/{to}/{rate} 
  
    http://localhost:8000/exchange/10/BRL/USD/4.50

    * Resposta:

      `code:200`

      ```json 
      {
        "valorConvertido": 45,
        "simboloMoeda": "$"
      }
      ```

    * Adicional:
        Caso não saiba o valor de cotação da moeda para enviar como o parametro **rate** na url, basta enviar o valor **0** que a partir da moeda eu faço uma integração com a [AwesomeApi](https://docs.awesomeapi.com.br/api-de-moedas) buscando essa cotação!

# 
* GetConsults:

    Listagem com os logs de todas as requisições de conversão feitas em GetExchange.

    http://localhost:8000/consults

* Resposta:

    `code:200`

    ```json
    [
	    {
	    	"id": 1,
	    	"amount": 10,
	    	"from": "BRL",
	    	"to": "USD",
	    	"rate": 4.5,
	    	"createdAt": "2023-07-13T07:16:13.802Z"
	    },
	    {
	    	"id": 2,
	    	"amount": 100,
	    	"from": "USD",
	    	"to": "BRL",
	    	"rate": 4.81895,
	    	"createdAt": "2023-07-13T07:16:34.113Z"
	    }
    ]
    ```