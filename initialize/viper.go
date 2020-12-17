package initialize

import (
	"ff/g"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitViper() {
	g.VP = viper.New()
	g.VP.SetConfigType("yaml")
	g.VP.AddConfigPath("./setting")
	// g.VP.AddConfigPath(".")

	if _, err := os.Stat("./setting/config.yaml"); err == nil {
		g.VP.SetConfigName("config")
	} else {
		g.VP.SetConfigName("config-example")
	}

	if err := g.VP.ReadInConfig(); err != nil {
		panic(fmt.Errorf("==>加载配置文件失败: %s \n", err))
	}
	fmt.Println("==>加载配置文件成功！")
}
