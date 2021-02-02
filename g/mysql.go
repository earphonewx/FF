package g

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

type mysql struct {
	*gorm.DB
	once sync.Once
}

var db mysql

func DB() *gorm.DB {
	db.once.Do(func() {
		var err error
		db.DB, err = gorm.Open("mysql",
			fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				VP.GetString("mysql.user"),
				VP.GetString("mysql.password"),
				VP.GetString("mysql.host"),
				VP.GetInt("mysql.port"),
				VP.GetString("mysql.db")))
		if err != nil {
			panic(fmt.Errorf("无法连接到MySQL: %s \n", err))
		}

		gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
			return defaultTableName
		}

		// 全局禁用表名复数
		db.DB.SingularTable(true)

		// 开启日志
		db.DB.LogMode(VP.GetBool("mysql.enable-sql-log"))

		// SetMaxIdleCons 设置连接池中的最大闲置连接数。
		db.DB.DB().SetMaxIdleConns(VP.GetInt("mysql.max-idle-conn"))

		// SetMaxOpenCons 设置数据库的最大连接数量。
		db.DB.DB().SetMaxOpenConns(VP.GetInt("mysql.max-open-conn"))

		// SetConnMaxLifetime 设置连接的最大可复用时间。
		db.DB.DB().SetConnMaxLifetime(time.Duration(VP.GetInt("mysql.conn-max-lifetime")) * time.Minute)

	})
	return db.DB
}
