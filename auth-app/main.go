package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func hello(ctx iris.Context) {
	ctx.JSON("Hello World!")
}

func main() {

	app := iris.New()
	app.Handle("GET", "/hello", hello)
	fmt.Println("Server Online!")
	app.Listen(":8080")

}