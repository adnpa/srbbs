package postgresql

import (
	"go.uber.org/zap"
	"golang.org/x/net/context"
	"srbbs/src/model"
	"srbbs/src/query"
	"srbbs/src/util"
	"time"
)

func CheckUserExist(username string) (bool, error) {
	u := query.User
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	cnt, err := query.User.WithContext(ctx).Where(u.Username.Eq(username)).Count()
	if err != nil {
		return false, err
	}
	return cnt > 0, nil
}

func CreateUser(user *model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	user.Password = util.EncryptPassword([]byte(user.Password))
	err := query.User.WithContext(ctx).Create(user)
	if err != nil {
		zap.L().Warn("error service sign in", zap.Error(err))
		return err
	}
	return nil
}

func GetUserByUserName(uname string) (*model.User, error) {
	u := query.User
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	first, err := query.User.WithContext(ctx).Where(u.Username.Eq(uname)).First()
	if err != nil {
		return nil, err
	}
	return first, err
}
