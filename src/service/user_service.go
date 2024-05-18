package service

import (
	"errors"
	"srbbs/src/dao/postgresql"
	"srbbs/src/enums"
	"srbbs/src/model"
	"srbbs/src/util"
	"srbbs/src/util/lib/algo/snowflake"
)

// SignUp 注册
func SignUp(form model.C2SRegister) (err error) {
	boo, err := postgresql.CheckUserExist(form.UserName)
	if err != nil {
		return err
	}
	if boo {
		return errors.New(enums.ErrorUserExit)
	}

	userId, err := snowflake.GetID()
	user := &model.User{
		UserID:   int64(userId),
		Username: form.UserName,
		Email:    form.Email,
		Password: form.Password,
	}
	err = postgresql.CreateUser(user)
	return
}

// LogIn 登录
func LogIn(form model.C2SLogIn) (user *model.User, err error) {
	user, err = postgresql.GetUserByUserName(form.UserName)
	if err != nil {
		return nil, errors.New(enums.ErrorUserNotExit)
	}
	formPwd := util.EncryptPassword([]byte(form.Password))

	if user.Password != formPwd {
		return nil, errors.New(enums.ErrorPasswordWrong)
	}
	return
}
