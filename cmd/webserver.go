package cmd

import (
	"ff/g"
	"ff/initialize"
	"ff/weather"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var S = &http.Server{}

var webServerCmd = &cobra.Command{
	Use:   "webserver",
	Short: "use this command to start ff web server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	webServerCmd.AddCommand(startWebServerCmd)
}

var startWebServerCmd = &cobra.Command{
	Use:   "start",
	Short: "use this command to start web server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		startWebServer()
	},
}

func startWebServer() {
	// 初始化路由
	rootRouter := initialize.InitRouter()

	// 程序结束后关闭数据库连接
	defer func() {
		if err := g.DB().Close(); err != nil {
			panic(fmt.Errorf("==>停止应用...关闭数据库连接时出错: %s \n", err))
		}
	}()

	// 定时爬取天气预报信息
	go weather.RealtimeWeather(time.Minute * g.VP.GetDuration("weather.update-frequency"))

	// HTTP配置
	S = &http.Server{
		Addr:           fmt.Sprintf(":%d", g.VP.GetInt("server.http-port")),
		Handler:        rootRouter,
		ReadTimeout:    g.VP.GetDuration("server.read-timeout") * time.Second,
		WriteTimeout:   g.VP.GetDuration("server.write-timeout") * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	g.Logger.Info("==================== 启动服务... ====================")
	if err := S.ListenAndServe(); err != nil {
		panic(fmt.Errorf("==>启动服务失败: %s \n", err))
	}
}
