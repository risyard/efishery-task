package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/risyard/efishery-task/fetch-app/cache"
	"github.com/risyard/efishery-task/fetch-app/config"
	"github.com/risyard/efishery-task/fetch-app/handler/komoditas"
	mw "github.com/risyard/efishery-task/fetch-app/middleware"
)

func hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Hello World!")
}

func main() {
	config.InitConfig()
	
	cacheWorker := cache.NewCacheWorker()
	cacheTimer := time.NewTicker(time.Duration(config.CacheDuration) * time.Minute)
	go cacheWorker.StartWorker(cacheTimer)
	
	app := gin.New()
	app.Use(mw.CheckJWTToken)

	app.GET("/hello", hello)

	komHandler := komoditas.NewKomoditasHandler()
	app.GET("/komoditas", komHandler.GetListKomoditas)

	listenPort := fmt.Sprintf(":%s", config.PORT)
	fmt.Println("Server online!")

	app.Run(listenPort)
}
