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
