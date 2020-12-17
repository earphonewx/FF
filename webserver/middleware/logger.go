package middleware

import (
	"ff/g"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
)

// GinLogger 接收gin框架默认的日志
func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		//start := time.Now()
		body, _ := ioutil.ReadAll(c.Request.Body)
		c.Next()
		//cost := time.Since(start)

		if g.VP.GetString("server.run-mode") == "debug" {
			g.Logger.Info(c.Request.URL.Path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.String("query", c.Request.URL.RawQuery),
				zap.String("body", string(body)),
				zap.String("client-ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				//zap.Duration("cost", cost),
			)
		}
	}
}
