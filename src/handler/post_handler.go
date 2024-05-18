package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"log"
	"srbbs/src/enums"
	"srbbs/src/model"
	"srbbs/src/service"
	"strconv"
)

//import (
//	"errors"
//	"github.com/gin-gonic/gin"
//	"github.com/go-playground/validator/v10"
//	"go.uber.org/zap"
//	"srbbs/src/dao/redis"
//	"srbbs/src/enums"
//	"srbbs/src/model"
//)

func CreatePostHandler(c *gin.Context) {
	var err error
	//  1 获取参数和校验数据有效性
	var post *model.Post
	if err = c.ShouldBindJSON(&post); err != nil {
		var errs validator.ValidationErrors
		if !errors.As(err, &errs) {
			ResponseError(c, enums.CodeInvalidParams)
			return
		}
		ResponseError(c, enums.CodeServerBusy)
		//ResponseErrorWithMsg(c, enums.CodeInvalidParams, removeTopStruct(errs.Translate(trans))) //注册参数错误
		return
	}

	//3获取自己的id
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, enums.CodeNotLogin)
		return
	}
	post.AuthorID = int64(userID)

	//4 调用service
	if err := service.CreatePost(strconv.FormatInt(post.AuthorID, 10), post); err != nil {
		ResponseError(c, enums.CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}

// 分页展示帖子列表
//func PostsHandler(c *gin.Context) {
//	// 获取分页参数
//	page, size := getPageInfo(c)
//	// 获取数据
//	data, err := service.GetPostList(page, size)
//	if err != nil {
//		ResponseError(c, enums.CodeServerBusy)
//		return
//	}
//	ResponseSuccess(c, data)
//}

// 根据社区id及时间或者分数排序分页展示帖子列表
func Posts2Handler(c *gin.Context) {
	// GET请求参数(query string)： /api/v1/posts2?page=1&size=10&order=time
	// 获取分页参数
	p := &model.C2SGetPostParam{}
	//c.ShouldBind() 根据请求的数据类型选择相应的方法去获取数据
	//c.ShouldBindJSON() 如果请求中携带的是json格式的数据，才能用这个方法获取到数据
	if err := c.ShouldBindQuery(p); err != nil {
		ResponseError(c, enums.CodeInvalidParams)
		return
	}

	// 获取数据
	data, err := service.GetPostListNew(p)
	if err != nil {
		ResponseError(c, enums.CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func PostDetailHandler(c *gin.Context) {
	// 1、获取参数(从URL中获取帖子的id)
	postIdStr := c.Param("id")
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		ResponseError(c, enums.CodeInvalidParams)
	}

	// 2、根据id取出id帖子数据(查数据库)
	post, err := service.GetPostById(postId)
	if err != nil {
		ResponseError(c, enums.CodeServerBusy)
	}

	// 3、返回响应
	ResponseSuccess(c, post)
}

func PostSearchHandler(c *gin.Context) {
	p := &model.C2SGetPostParam{}
	if err := c.ShouldBindQuery(p); err != nil {
		ResponseError(c, enums.CodeInvalidParams)
		return
	}
	// 获取数据
	data, err := service.PostSearch(p)
	if err != nil {
		ResponseError(c, enums.CodeServerBusy)
		return
	}

	log.Println("PostSearchHandler", data)
	ResponseSuccess(c, data)
}
