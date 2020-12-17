package cmd

import (
	"ff/g"
	"ff/initialize"
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

	// 初始化数据库配置
	initialize.InitMysql()

	// 初始化http客户端连接池
	initialize.InitHttpClient()

	// 程序结束后关闭数据库连接
	defer func() {
		if err := g.DB.Close(); err != nil{
			panic(fmt.Errorf("==>停止应用...关闭数据库连接时出错: %s \n", err))
		}
	}()

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
