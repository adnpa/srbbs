package test

import (
	"testing"
)

func TestSignInHandler(t *testing.T) {
	//// 创建 Gin 路由器
	////r := gin.Default()
	//r := gin.New()
	////r.POST("/api/v1/signup", handler.SignInHandler)
	//
	//r.Use(func(c *gin.Context) {
	//	start := time.Now()
	//	c.Next()
	//	duration := time.Since(start)
	//	t.Log("request processed",
	//		zap.String("method=", c.Request.Method),
	//		zap.String("path=", c.Request.URL.Path),
	//		zap.String("handler=", c.HandlerName()),
	//		zap.Int("status=", c.Writer.Status()),
	//		zap.Duration("duration=", duration),
	//	)
	//})
	//
	//srserver.InitServer(r)
	//
	//// 构建 POST 请求
	//formData := url.Values{
	//	"username":         {"test_user"},
	//	"email":            {"test_mail"},
	//	"gender":           {"test_gender"},
	//	"password":         {"test_password"},
	//	"confirm_password": {"test_confirm_password"},
	//}
	//req, _ := http.NewRequest("POST", "/api/v1/signup", strings.NewReader(formData.Encode()))
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//
	//// 创建响应记录器
	//w := httptest.NewRecorder()
	//
	//// 处理请求
	//r.ServeHTTP(w, req)
	//
	//// 断言响应
	//t.Log("code: ", w.Code)
	//assert.Equal(t, http.StatusOK, w.Code)
	////assert.Contains(t, w.Body.String(), "Registered successfully")
	//
	//time.Sleep(1 * time.Second)
}

//func registerHandler(c *gin.Context) {
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//
//	// 在这里进行注册逻辑
//	fmt.Printf("Registered user: %s/%s\n", username, password)
//	c.String(http.StatusOK, "Registered successfully")
//}
