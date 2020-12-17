package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Pagination struct {
	Page     int `form:"page" binding:"min=0,max=10000"`
	PageSize int `form:"pageSize" binding:"gte=0,lte=1000"`
}

func PaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagination Pagination
		if err := c.BindQuery(&pagination); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else if pagination.PageSize == 0 || pagination.Page == 0 {
			// 如果没有传分页参数，默认第一页取十条
			pagination.Page = 1
			pagination.PageSize = 10
		}
		// 设置分页变量
		c.Set("page_size", pagination.PageSize)
		c.Set("page", pagination.Page)

		c.Next()
	}
}
