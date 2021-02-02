package cmd

import (
	"ff/g"
	"ff/model"
	"fmt"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "use this command to migrate your db",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("==>开始迁移数据库...")
		// 初始化数据库
		//initialize.InitMysql()
		// 迁移结束关闭数据库
		defer func() {
			if err := g.DB().Close(); err != nil {
				panic(fmt.Errorf("==>任务结束...关闭数据库连接时出错: %s \n", err))
			}
		}()
		g.DB().AutoMigrate(&model.Bookmark{}, &model.Memo{}, &model.AuthUser{})
		fmt.Println("==>数据库迁移完毕!")
	},
}
