package main

import (
	"ff/cmd"
	"fmt"
)

// @title ff API
// @version latest
// @description This is ff.

// @contact.name fengfeng
// @contact.url http://www.earphonewx.top
// @contact.email earphonewx@163.com

// @host localhost:8888
// @BasePath /api
func main() {
	if err := cmd.Execute(); err != nil {
		panic(fmt.Errorf("==>初始化命令行工具异常: %s \n", err))
	}
}
