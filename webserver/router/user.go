package router

import (
	"ff/webserver/api"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("")
	{
		//bookmarkRouter.GET("/user", api.GetBookmark)
		// 用户注册
		userRouter.POST("/user", api.AddBookmark)
		// 修改书签
		//bookmarkRouter.PATCH("/bookmark/:id", api.EditBookmark)
		// 删除书签
		//bookmarkRouter.DELETE("/bookmark/:id", api.DeleteBookmark)
	}
}
