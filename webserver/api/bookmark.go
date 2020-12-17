package api

import (
	"ff/model"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
	"net/http"
	"strconv"
)

type Bookmark struct {
	Id    uint   `form:"id" json:"id" binding:"min=0"`
	Title string `form:"title" json:"title" binding:"max=20"`
	Url   string `form:"url" json:"url" binding:"max=100,min=0"`
	Type  string `form:"type" json:"type" binding:"max=20"`
}

// @Summary
// @Description 可以根据id、title、url、type过滤获取相应书签列表
// @Tags 获取书签
// @Produce  json
// @Param id query int false "ID"
// @Param title query string false "Title"
// @Param url query string false "Url"
// @Param type query string false "Type"
// @Success 200 {string} json "{"current_page":1,"count":100,"data":[]}"
// @Router /bookmark [get]
func GetBookmark(c *gin.Context) {
	var params Bookmark
	// 验证查询参数
	if err := c.BindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// 取数据
	offset := (c.GetInt("page") - 1) * c.GetInt("page_size")
	data, err := model.GetBookmark(offset, c.GetInt("page_size"), &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	// 取count
	count, countErr := model.GetBookmarkCount(&params)
	if countErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": countErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"page": c.GetInt("page"),
		"count":        count,
		"data":         data})
}

func AddBookmark(c *gin.Context) {
	var bookmark Bookmark
	// 验证post数据
	if err := c.BindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// url不能为空
	if bookmark.Url == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "url can not be blank"})
		return
	}

	// 返回结构和GetBookmark接口的验证共用了一个结构体，改动需注意
	if err := model.AddBookmark(&bookmark); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": bookmark, "msg": "success"})
}

func EditBookmark(c *gin.Context) {
	// 取对应的id并检验
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid id"})
		return
	}

	var bookmark Bookmark
	// 验证修改的数据
	if err := c.BindJSON(&bookmark); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// 修改
	if err := model.EditBookmark(uint(id), &bookmark); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "success"})
}

func DeleteBookmark(c *gin.Context) {
	// 取对应的id并校验
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid id"})
		return
	}

	// 删除
	if err := model.DeleteBookmark(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"msg": "success"})
}
