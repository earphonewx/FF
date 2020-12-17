package test

import (
	"bytes"
	"encoding/json"
	"ff/initialize"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookmark(t *testing.T) {
	// 初始化路由
	r := initialize.InitRouter()
	// 初始化数据库
	initialize.InitMysql()

	// 获取书签列表
	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/api/bookmark", nil)
	r.ServeHTTP(w1, req1)
	assert.Equal(t, 200, w1.Code)

	// 测试分页
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/api/bookmark?page=10&pageSize=89", nil)
	r.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)

	// 新建书签
	w3 := httptest.NewRecorder()
	b3 := map[string]interface{}{
		"url": "0123456789876543210.com",
	}
	jsonByte, _ := json.Marshal(b3)
	req3, _ := http.NewRequest("POST", "/api/bookmark", bytes.NewReader(jsonByte))
	r.ServeHTTP(w3, req3)
	assert.Equal(t, 201, w3.Code)

	// 保存一下创建的书签的id，留着后面patch和delete用
	resp := w3.Body.Bytes()
	var res map[string]map[string]interface{}
	if err := json.Unmarshal(resp, &res); err != nil {
		panic(err)
	}
	fmt.Println("delete bookmark:", res)
	id := int(res["data"]["id"].(float64))

	// 修改书签
	w4 := httptest.NewRecorder()
	b4 := map[string]interface{}{
		"title": "0123456789876543210",
	}
	patchData, _ := json.Marshal(b4)
	req4, _ := http.NewRequest("PATCH", "/api/bookmark/"+strconv.Itoa(id), bytes.NewReader(patchData))
	r.ServeHTTP(w4, req4)
	assert.Equal(t, 200, w4.Code)

	// 删除书签
	w5 := httptest.NewRecorder()
	req5, _ := http.NewRequest("DELETE", "/api/bookmark/"+strconv.Itoa(id), nil)
	r.ServeHTTP(w5, req5)
	assert.Equal(t, 204, w5.Code)
}
