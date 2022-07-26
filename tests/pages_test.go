package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHomePage(t *testing.T) {

	// 参考：https://learnku.com/courses/go-basic/1.17/refactoring-testing/11517
	baseURL := "http://localhost:3000"
	// 1. 请求，模拟用户访问浏览器
	var (
		resp *http.Response
		err  error
	)

	//Go 的 http 包兼具 HTTP 服务器和 HTTP 客户端的功能，HTTP 客户端支持 GET/POST/PUT 等请求方式，常用于访问网页，或者请求第三方 API
	resp, err = http.Get(baseURL + "/")
	// 2. 检测 —— 是否无错误且 200
	assert.NoError(t, err, "有错误发生，err 不为空")
	assert.Equal(t, 200, resp.StatusCode, "应返回状态码 200")
}
