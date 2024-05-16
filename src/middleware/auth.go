package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"srbbs/src/enums"
	"srbbs/src/handler"
	"srbbs/src/util/jwt"
)

// JWTAuthMiddleware 基于jwt的授权中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		aToken, _ := handler.ParseJwtHeader(c)
		// 处理
		claims, err := jwt.ParseToken(aToken)
		if err != nil {
			fmt.Println(err)
			handler.ResponseError(c, enums.CodeInvalidToken)
			c.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文c上
		c.Set(handler.ContextUserIDKey, claims.UserID)
		c.Next() // 后续的处理函数可以用过c.Get(ContextUserIDKey)来获取当前请求的用户信息
	}
}
