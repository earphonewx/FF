package initialize

import (
	"ff/g"
	_ "ff/webserver/docs"
	"ff/webserver/middleware"
	"ff/webserver/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	gin.SetMode(g.VP.GetString("server.run-mode"))
	myRouter := gin.New()
	myRouter.Use(middleware.ZapLogger())
	myRouter.Use(middleware.CustomSimpleRecovery())
	myRouter.Use(middleware.PaginationMiddleware())
	//myRouter.Use(middleware.JWTAuthMiddleware())


	ApiGroup := myRouter.Group("/api")

	// swagger api文档
	myRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 注册路由
	router.InitBookmarkRouter(ApiGroup)
	router.InitMemoRouter(ApiGroup)
	router.InitWeatherRouter(ApiGroup)

	return myRouter
}
