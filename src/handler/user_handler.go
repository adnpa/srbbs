package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"srbbs/src/enums"
	"srbbs/src/model"
	"srbbs/src/service"
	"srbbs/src/srlogger"
	"srbbs/src/util/jwt"
	"strings"
)

// SignUpHandler 注册
func SignUpHandler(c *gin.Context) {
	var err error
	//  1 获取参数
	var form model.RegisterForm
	//  2 校验数据有效性
	if err = c.ShouldBindJSON(&form); err != nil {
		//zap.L().Error("sign up with invalid param", zap.Error(err), zap.Any("form param", form))
		var errs validator.ValidationErrors
		if !errors.As(err, &errs) {
			zap.L().Info("form param")
			ResponseError(c, enums.CodeInvalidParams)
			return
		}
		ResponseError(c, enums.CodeInvalidParams)
		//ResponseErrorWithMsg(c, enums.CodeInvalidParams, removeTopStruct(errs.Translate(trans))) //注册参数错误
		return
	}

	//  3 业务逻辑
	if err = service.SignUp(form); err != nil {
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("logic.signup failed", zap.Error(err))
		if err.Error() == enums.ErrorUserExit {
			ResponseError(c, enums.CodeUserExist)
			return
		}
		ResponseError(c, enums.CodeServerBusy)
		return
	}

	//	4 返回
	srlogger.Logger().Info("sign up success")
	ResponseSuccess(c, nil)

}

func LogInHandler(c *gin.Context) {
	var err error
	//  1 获取参数
	var form model.LogInForm
	//  2 校验数据有效性
	if err = c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		zap.L().Error("log in with invalid param", zap.Error(err), zap.Any("form param", form))
		var errs validator.ValidationErrors
		if !errors.As(err, &errs) {
			zap.L().Info("form param")
			ResponseError(c, enums.CodeInvalidParams)
			return
		}
		ResponseError(c, enums.CodeInvalidParams)
		return
	}
	//	3 业务逻辑
	user, err := service.LogIn(form)
	if err != nil {
		zap.L().Error("logic.Login failed", zap.String("username", form.UserName), zap.Error(err))
		if err.Error() == enums.ErrorUserNotExit {
			ResponseError(c, enums.CodeUserNotExist)
			return
		}
		ResponseError(c, enums.CodeInvalidParams)
		return
	}
	aToken, rToken, err := jwt.GenToken(uint64(user.UserID), user.Username)
	ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", user.UserID), //js识别的最大值：id值大于1<<53-1  int64: i<<63-1
		"user_name":     user.Username,
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}

func RefreshTokenHandler(c *gin.Context) {
	oAToken, oRToken := ParseJwtHeader(c)
	aToken, rToken, err := jwt.RefreshToken(oAToken, oRToken)
	if err != nil {
		ResponseError(c, enums.CodeInvalidToken)
	}
	//if aToken == "" && rToken == "" && err == nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"access_token":  oAToken,
	//		"refresh_token": oRToken,
	//	})
	//	return
	//}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}

// ParseJwtHeader 辅助函数
func ParseJwtHeader(c *gin.Context) (aToken, rToken string) {
	//rt := c.Query("refresh_token")
	// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
	// 这里假设Token放在Header的 Authorization 中，并使用 Bearer 开头
	// 这里的具体实现方式要依据你的实际业务情况决定
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, enums.CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}
	//	空格分割aToken和rToken
	//parts := strings.SplitN(authHeader, " ", 2)
	parts := strings.SplitN(authHeader, " ", 3)
	if !(len(parts) == 3 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, enums.CodeInvalidToken, "Token格式不对")
		c.Abort()
		return
	}
	aToken = parts[1]
	rToken = parts[2]
	return
}
