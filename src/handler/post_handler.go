package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"srbbs/src/dao/redis"
	"srbbs/src/enums"
	"srbbs/src/model"
)

func CreatePostHandler(c *gin.Context) {
	var err error
	//  1 获取参数
	var post model.Post
	//  2 校验数据有效性
	if err = c.ShouldBindJSON(&post); err != nil {
		var errs validator.ValidationErrors
		if !errors.As(err, &errs) {
			zap.L().Info("form param")
			ResponseError(c, enums.CodeInvalidParams)
			return
		}
		ResponseError(c, enums.CodeServerBusy)
		//ResponseErrorWithMsg(c, enums.CodeInvalidParams, removeTopStruct(errs.Translate(trans))) //注册参数错误
		return
	}

	//3 创建帖子
	if err:= redis.Close()


	ResponseSuccess(c, nil)
}
