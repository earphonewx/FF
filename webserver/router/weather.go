package router

import (
	"ff/webserver/api"
	"github.com/gin-gonic/gin"
)

func InitWeatherRouter(Router *gin.RouterGroup) {
	weatherRouter := Router.Group("")
	{
		// 查询城市编码
		weatherRouter.GET("/weather/cityCode", api.SearchCity)
		// 查询天气情况
		weatherRouter.GET("/weather", api.CityWeather)
	}
}