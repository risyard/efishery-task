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

	listenPort := fmt.Sprintf(":%s", config.PORT)
	fmt.Println("Server online!")
	
	app.Listen(listenPort)

}
