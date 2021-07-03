package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/risyard/efishery-task/fetch-app/config"
	"github.com/risyard/efishery-task/fetch-app/handler/komoditas"
)

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}

func main() {
	config.InitConfig()
	app := gin.Default()

	app.GET("/hello", hello)

	komHandler := komoditas.NewKomoditasHandler()
	app.GET("/komoditas", komHandler.GetListKomoditas)

	listenPort := fmt.Sprintf(":%s", config.PORT)
	fmt.Println("Server online!")

	app.Run(listenPort)
}
