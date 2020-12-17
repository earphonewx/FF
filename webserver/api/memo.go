package api

import (
	"ff/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Memo struct {
	ID       uint      `form:"id" json:"id" binding:"min=0"`
	Content  string    `form:"content" json:"content" binding:"max=100"`
	Deadline time.Time `form:"deadline" json:"deadline" binding:"omitempty"`
}

func GetMemo(c *gin.Context) {
	var params Memo
	// 验证查询参数
	if err := c.BindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 取数据
	offset := (c.GetInt("page") - 1) * c.GetInt("page_size")
	data, err := model.GetMemo(offset, c.GetInt("page_size"), &params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 取count
	count, countErr := model.GetMemoCount(&params)
	if countErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": countErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"current_page": c.GetInt("page"),
		"count":        count,
		"data":         data})
}

func AddMemo(c *gin.Context) {
	var memo Memo
	// 验证post数据
	if err := c.BindJSON(&memo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// content不能为空
	if memo.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "content can not be blank"})
		return
	}

	if err := model.AddMemo(&memo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": memo})
}

func EditMemo(c *gin.Context) {
	// 取对应的id并校验
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var memo Memo
	// 验证修改的数据
	if err := c.BindJSON(&memo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 修改
	if err := model.EditMemo(uint(id), &memo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "update success"})
}

func DeleteMemo(c *gin.Context) {
	// 取对应的id并校验
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	// 删除
	if err := model.DeleteMemo(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
