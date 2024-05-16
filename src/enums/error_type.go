package enums

import "errors"

// 数据库错误
var (
	ErrorUserExit      = "用户已存在"
	ErrorUserNotExit   = "用户不已存在"
	ErrorPasswordWrong = "密码错误"
	ErrorGenIDFailed   = errors.New("创建用户ID失败")
	ErrorInvalidID     = "无效的ID"
	ErrorQueryFailed   = "查询数据失败"
	ErrorInsertFailed  = errors.New("插入数据失败")
)
