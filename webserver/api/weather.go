package api

import (
	"ff/weather"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SearchCity(c *gin.Context) {
	addr := c.Query("addr")
	if addr == "" {
		c.JSON(http.StatusOK, gin.H{"data": [0]int{}})
		return
	}
	res := weather.QueryCity(addr)
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func CityWeather(c *gin.Context) {
	adcode := c.Query("adcode")
	if adcode == "" {
		c.JSON(http.StatusOK, gin.H{"data": [0]int{}})
		return
	}
	// 先从缓存拿
	if res, err := weather.WeatherCached(adcode); (err == nil) && (res != nil) && (len(res) != 0) {
		c.JSON(http.StatusOK, gin.H{"data": res})
		return
	}
	// 缓存没有就实时拉取并更新缓存
	weatherNow := weather.WeatherInfoNow(adcode)
	weather.CacheWeather(weatherNow)
	c.JSON(http.StatusOK, gin.H{"data": weatherNow})
}