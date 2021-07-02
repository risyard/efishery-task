package main

import (
	"fmt"

	"github.com/efishery-task/auth-app/handler"
	"github.com/efishery-task/auth-app/config"
	"github.com/kataras/iris/v12"
)

func main() {

	config.InitConfig()

	app := iris.New()
	h := handler.NewHandler()

	app.Handle("GET", "/hello", h.Hello)
	app.Handle("POST", "/user", h.AddUser)
	app.Handle("GET", "/token", h.GetToken)
	app.Handle("GET", "/claims", h.GetClaims)

	listenPort := fmt.Sprintf(":%s", config.PORT)
	fmt.Println("Server online!")
	
	app.Listen(listenPort)

}
