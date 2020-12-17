package router

import (
	"ff/webserver/api"
	"ff/webserver/middleware"
	"github.com/gin-gonic/gin"
)

func InitMemoRouter(Router *gin.RouterGroup) {
	memoRouter := Router.Group("").Use(middleware.PaginationMiddleware())
	{
		// 获取备忘录
		memoRouter.GET("/memo", api.GetMemo)
		// 新建备忘录
		memoRouter.POST("/memo", api.AddMemo)
		// 更新备忘录
		memoRouter.PATCH("/memo/:id", api.EditMemo)
		// 删除备忘录
		memoRouter.DELETE("/memo/:id", api.DeleteMemo)
	}
}
