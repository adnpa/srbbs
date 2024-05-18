package model

// C2SRegister 注册请求参数
type C2SRegister struct {
	UserName        string `json:"username" binding:"required"`  // 用户名
	Email           string `json:"email" binding:"required"`     // 邮箱
	Gender          int    `json:"gender" binding:"oneof=0 1 2"` // 性别 0:未知 1:男 2:女
	Password        string `json:"password" binding:"required"`  // 密码
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type C2SLogIn struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
