package main

import (
	"github.com/gin-gonic/gin"
	"./conf"
	"./middleware"
	"./model"
)

func main() {
	if err := conf.InitConfig(); err != nil {
		panic(err)
	}

	// set mode
	gin.SetMode(gin.ReleaseMode)

	// start server
	r := gin.New()

	// add middleware
	r.Use(middleware.Logger())

	// router
	r.GET("/", model.HandleCheckSignature)
	r.POST("/", model.HandleRequest)
	r.Run(conf.Conf.App.Port)
}
