package test

import (
	"bytes"
	"encoding/json"
	"ff/initialize"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMemo(t *testing.T) {
	// 初始化路由
	r := initialize.InitRouter()
	// 初始化数据库
	initialize.InitMysql()

	// GET
	w1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("GET", "/api/memo", nil)
	r.ServeHTTP(w1, req1)
	assert.Equal(t, 200, w1.Code)

	// 测试分页
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/api/memo?page=10&pageSize=89", nil)
	r.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)

	// 新建
	w3 := httptest.NewRecorder()
	b3 := map[string]interface{}{
		"content":  "0123456789876543210",
		"deadline": time.Now(),
	}
	jsonByte, _ := json.Marshal(b3)
	req3, _ := http.NewRequest("POST", "/api/memo", bytes.NewReader(jsonByte))
	r.ServeHTTP(w3, req3)
	assert.Equal(t, 201, w3.Code)

	// 保存一下id，留着后面patch和delete用
	resp := w3.Body.Bytes()
	var res map[string]map[string]interface{}
	if err := json.Unmarshal(resp, &res); err != nil {
		panic(err)
	}
	id := int(res["data"]["id"].(float64))

	// 修改
	w4 := httptest.NewRecorder()
	b4 := map[string]interface{}{
		"content": "meouwcvytwewiinowhiu",
	}
	patchData, _ := json.Marshal(b4)
	req4, _ := http.NewRequest("PATCH", "/api/memo/"+strconv.Itoa(id), bytes.NewReader(patchData))
	r.ServeHTTP(w4, req4)
	assert.Equal(t, 200, w4.Code)

	// 删除
	w5 := httptest.NewRecorder()
	req5, _ := http.NewRequest("DELETE", "/api/memo/"+strconv.Itoa(id), nil)
	r.ServeHTTP(w5, req5)
	assert.Equal(t, 204, w5.Code)
}
