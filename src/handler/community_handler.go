package handler

import (
	"github.com/gin-gonic/gin"
	"srbbs/src/enums"
	"srbbs/src/service"
	"strconv"
)

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区(community_id,community_name)以列表的形式返回
	communityList, err := service.GetCommunityList()
	if err != nil {
		ResponseError(c, enums.CodeServerBusy) // 不轻易把服务端报错暴露给外面
		return
	}
	ResponseSuccess(c, communityList)
}

func CommunityDetailHandler(c *gin.Context) {
	// 1、获取社区ID
	communityIdStr := c.Param("id")                               // 获取URL参数
	communityId, err := strconv.ParseUint(communityIdStr, 10, 64) // id字符串格式转换
	if err != nil {
		ResponseError(c, enums.CodeInvalidParams)
		return
	}

	// 2、根据ID获取社区详情
	communityList, err := service.GetCommunityDetailByID(int(communityId))
	if err != nil {
		ResponseErrorWithMsg(c, enums.CodeSuccess, err.Error())
		return
	}
	ResponseSuccess(c, communityList)
}
