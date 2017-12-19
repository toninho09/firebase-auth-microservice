# Api de Autenticação

Api utilizada para validar a autenticação dos usuários e obter as informações adicionais do firebase

## Requisitos

* Go >= 1.9.2

#### Instalando dependencias:

````shell
    go get -u firebase.google.com/go
    go get -u github.com/gin-gonic/gin
````

## Build


````shell
    go build -o api .
````



## Configurações

As configurações são feitas utilizando variáveis de ambiente 

* PORT : Configura a porta que o serviço estará disponível, default  `8080` 
* ENV : Define o ambiente em que o sistema está rodando, Valores `DEV` , `PROD`

## Docker

Para criar a imagem da aplicação utilizar o comando

````shell
docker build -t auth .
````
 
# Serviços


### [POST] /verify

Verifica se o token e válido, caso o token não seja valido, o HTTP_STATUS de retorno será 401

##### request

````json
{
	"token":"eyJhbGciOiJSUzI1NiIsImt......"	
}
````

##### response

````json
{
	"message": "Ok",
	"data": null
}
````

### [POST] /get-user

Obtém as informações do usuário pelo firebase

##### Request

````json
{
	"token":"eyJhbGciOiJSUzI1NiIsImt......"	
}
````

##### Response

````json
{
	"displayName": "Display name",
	"email": "email@email.com",
	"photoUrl": "https://.../photo.jpg",
	"localId": "localId",
	"CustomClaims": null,
	"Disabled": false,
	"EmailVerified": true,
	"ProviderUserInfo": [
		{
			"displayName": "Display name",
			"email": "email@email.com",
			"photoUrl": "https://.../photo.jpg",
			"providerId": "google.com"
		}
	],
	"UserMetadata": {
		"CreationTimestamp": 1513099763000,
		"LastLogInTimestamp": 1513100854000
	}
}
````

