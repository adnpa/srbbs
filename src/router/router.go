package router

import (
	"github.com/gin-gonic/gin"
	"srbbs/src/handler"
	"srbbs/src/middleware"
)

// 添加路由

func SetUpRouter(r *gin.Engine, mode string) {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//注册登录
	v1 := r.Group("/api/v1")
	v1.POST("/signup", handler.SignUpHandler)
	v1.POST("/login", handler.LogInHandler)
	v1.GET("/refresh_token", handler.RefreshTokenHandler) // 刷新accessToken

	//帖子
	//v1.GET("posts", handler.PostsHandler)
	v1.GET("/post", handler.Posts2Handler)
	v1.GET("/post/:id", handler.PostDetailHandler)
	v1.GET("/search", handler.PostSearchHandler)

	//社区
	v1.GET("/community", handler.CommunityHandler)
	v1.GET("/community/:id", handler.CommunityDetailHandler)

	//授权中间件
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.POST("/post", handler.CreatePostHandler)
		//v1.POST("/community", handler.CommunityHandler) // 获取分类社区列表

		//v1.POST()
	}

	//性能测试
	//pprof.Register(v1)
	// 404
	//r.NoRoute(func(c *gin.Context) {
	//	log.Println("没有找到页面")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "404",
	//	})
	//})
}
