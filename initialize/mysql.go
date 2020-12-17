package initialize

import (
	"ff/g"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func InitMysql() {
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			g.VP.GetString("mysql.user"),
			g.VP.GetString("mysql.password"),
			g.VP.GetString("mysql.host"),
			g.VP.GetInt("mysql.port"),
			g.VP.GetString("mysql.db-name")))
	if err != nil {
		panic(fmt.Errorf("==>初始化MySQL配置失败: %s \n", err))
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}

	// 全局禁用表名复数
	db.SingularTable(true)

	// 开启日志
	db.LogMode(g.VP.GetBool("mysql.enable-sql-log"))

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(g.VP.GetInt("mysql.max-idle-conn"))

	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(g.VP.GetInt("mysql.max-open-conn"))

	// SetConnMaxLifetime 设置连接的最大可复用时间。
	db.DB().SetConnMaxLifetime(time.Duration(g.VP.GetInt("mysql.conn-max-lifetime")) * time.Minute)

	g.DB = db
}
