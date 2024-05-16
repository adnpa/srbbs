package main

import (
	//匿名导入 初始化各类资源
	_ "srbbs/src/conf"
	_ "srbbs/src/dao/postgresql"
	_ "srbbs/src/dao/redis"
	_ "srbbs/src/srlogger"
	_ "srbbs/src/util/lib/algo/snowflake"

	"srbbs/src/srserver"
)

func main() {
	server := srserver.NewServer()
	server.Init()
	server.Start()

	//r := gin.New()
	//
	//r.Use(func(c *gin.Context) {
	//	start := time.Now()
	//	log.Println("method=", c.Request.Method, "path=", c.Request.URL.Path, "status=", c.Writer.Status())
	//	//log.Println(c.GetRawData())
	//	//log.Println(c.Params)
	//	//log.Println(c.HandlerName())
	//	c.Next()
	//	duration := time.Since(start)
	//	log.Println("method=", c.Request.Method, "path=", c.Request.URL.Path, "status=", c.Writer.Status(), "duration=", duration)
	//})
	//
	//r.GET("/hello", func(c *gin.Context) {
	//	log.Println("hello")
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "welcome to rsbbs",
	//	})
	//})
	//
	//r.POST("/hello", func(c *gin.Context) {
	//	rspMap := map[string]interface{}{ // 也可用结构体方式返回
	//		"code": 0,
	//		"rsp":  fmt.Sprintf("welcome to rsbbs"),
	//	}
	//
	//	c.JSON(http.StatusOK, rspMap)
	//	fmt.Printf("rsp: %+v\n", rspMap)
	//
	//})
}
