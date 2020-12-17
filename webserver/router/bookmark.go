package router

import (
	"ff/webserver/api"
	"github.com/gin-gonic/gin"
)

func InitBookmarkRouter(Router *gin.RouterGroup) {
	bookmarkRouter := Router.Group("")
	{
		// 获取书签
		bookmarkRouter.GET("/bookmark", api.GetBookmark)
		// 创建书签
		bookmarkRouter.POST("/bookmark", api.AddBookmark)
		// 修改书签
		bookmarkRouter.PATCH("/bookmark/:id", api.EditBookmark)
		// 删除书签
		bookmarkRouter.DELETE("/bookmark/:id", api.DeleteBookmark)
	}
}
