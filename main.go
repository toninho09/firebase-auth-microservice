package main

import (
	"golang.org/x/net/context"
	"fmt"
	"firebase.google.com/go"
	"google.golang.org/api/option"
	"github.com/gin-gonic/gin"
	"firebase.google.com/go/auth"
)

var client *auth.Client
var ctx context.Context

func main() {
	var err error
	ctx = context.Background()
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("./api.json"))

	if err != nil {
		fmt.Errorf("error getting Auth client: %v", err)
		panic(err)
	}

	client, err = app.Auth(ctx)
	if err != nil {
		fmt.Errorf("error getting Auth client: %v", err)
		panic(err)
	}

	if getEnvironmentRelease() == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	router.POST("/verify", verify)
	router.POST("/get-user", getUser)

	router.Run(":" + getPort())

}

func verify(c *gin.Context) {

	var json VerifyStruct
	if err := c.ShouldBindJSON(&json); err == nil {

		_, err := client.VerifyIDToken(json.Token)
		if err != nil {
			c.JSON(401, returnResponse( ERROR_TOKEN_INVALID))
			panic(err)
		}
		c.JSON(200, returnResponse( MESSAGE_OK))
	} else {
	c.JSON(500, returnResponse( ERROR_MESSAGE_GENERIC))
	panic(err)
	}
}

func getUser(c *gin.Context) {

	var json VerifyStruct
	if err := c.ShouldBindJSON(&json); err == nil {

		token, err := client.VerifyIDToken(json.Token)
		if err != nil {
			c.JSON(401, returnResponse( ERROR_TOKEN_INVALID))
			panic(err)
		}

		user, err := client.GetUser(ctx, token.UID)

		if err != nil {
			c.JSON(500, returnResponse( ERROR_CANT_GET_USER))
			panic(err)
		}
		c.JSON(200, user)
	} else {
		c.JSON(500, returnResponse( ERROR_MESSAGE_GENERIC))
		panic(err)
	}
}

