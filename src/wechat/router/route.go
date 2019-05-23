package router

import (
	"github.com/gin-gonic/gin"
	"wechat/api"
	//"wechat/middleware"
)

func InitRouter() *gin.Engine {

	// set mode
	gin.SetMode(gin.ReleaseMode)
	// start server
	//r := gin.New()
	// add middleware
	//r.Use(middleware.Logger())
	r := gin.Default()

	r.GET("/", api.HandleCheckSignature)
	r.POST("/", api.HandleRequest)

	//testRoute := r.Group("/test"){
	//	testRoute.get("/", )
	//	testRout.POST("/data", )
	//
	//}

	return r
}