package service_test

import (
	"github.com/stretchr/testify/suite"
	"srbbs/src/conf"
	"srbbs/src/dao/postgresql"
	"srbbs/src/util"
	"srbbs/src/util/snowflake"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
}

func (suite *UserTestSuite) SetupTest() {
	suite.Nil(conf.InitCfg())
	suite.Nil(postgresql.Init(conf.Cfg.PostgresConfig))
	suite.Nil(snowflake.Init(conf.Cfg.MachineId))
}

//func (suite *UserTestSuite) TestCreateUser() {
//	userId, _ := snowflake.GetID()
//	u := &model.User{
//		ID:       1,
//		Username: "user1",
//		UserID:   int64(userId),
//		Email:    "12345@qq.com",
//	}
//	err := postgresql.CreateUser(u)
//	suite.Nil(err)
//}
//
//func (suite *UserTestSuite) TestCheckUserExist() {
//	boo, err := postgresql.CheckUserExist("user1")
//	log.Println(boo, err)
//	//suite.True(boo)
//	suite.False(boo)
//	suite.Nil(err)
//}

func (suite *UserTestSuite) TestEncryptPassword() {
	util.EncryptPassword([]byte("abcabc"))
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
