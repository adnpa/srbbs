package test

import (
	"github.com/stretchr/testify/suite"
	"log"
	"srbbs/src/util/jwt"
	"testing"
)

type JWTTestSuite struct {
	suite.Suite
}

func (suite *JWTTestSuite) SetupTest() {
	//suite.Nil(conf.InitCfg())
}

func (suite *JWTTestSuite) TestGenToken() {
	aToken, rToken, err := jwt.GenToken(123, "abc")
	log.Println("TestGenToken", aToken)
	log.Println("TestGenToken", rToken)
	suite.Nil(err)
}

func (suite *JWTTestSuite) TestParseToken() {
	aToken, _, err := jwt.GenToken(123, "abc")
	claims, err := jwt.ParseToken(aToken)
	log.Println("TestParseToken", claims)
	suite.Nil(err)
}

func (suite *JWTTestSuite) TestRefreshToken() {
	aToken, rToken, err := jwt.GenToken(123, "abc")
	nAToken, nRToken, err := jwt.RefreshToken(aToken, rToken)
	log.Println("TestRefreshToken old aToken: ", aToken)
	log.Println("TestRefreshToken old rToken: ", rToken)

	// 没有过期这里为nil
	log.Println("TestRefreshToken new aToken: ", nAToken)
	log.Println("TestRefreshToken new rToken: ", nRToken)

	suite.Nil(err)
}

func TestJWTTestSuite(t *testing.T) {
	suite.Run(t, new(JWTTestSuite))
}
