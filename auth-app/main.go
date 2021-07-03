package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/risyard/efishery-task/auth-app/config"
	"github.com/risyard/efishery-task/auth-app/handler/token"
	"github.com/risyard/efishery-task/auth-app/handler/user"
)

func main() {

	config.InitConfig()

	app := iris.New()

	userHandler := user.NewUserHandler()
	app.Handle("GET", "/hello", userHandler.Hello)
	app.Handle("POST", "/user", userHandler.AddUser)

	tokenHandler := token.NewTokenHandler()
	app.Handle("GET", "/token", tokenHandler.GetToken)
	app.Handle("GET", "/claims", tokenHandler.GetClaims)

	listenPort := fmt.Sprintf(":%s", config.PORT)
	fmt.Println("Server online!")
	
	app.Listen(listenPort)

}
