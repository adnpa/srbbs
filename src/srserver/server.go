package srserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"srbbs/src/conf"
	"srbbs/src/router"
	"srbbs/src/srlogger"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Init() {
	r := gin.New()
	s.Engine = r

	// read config
	//var err error
	//if err = conf.InitCfg(); err != nil {
	//	zap.L().Warn("error reading config", zap.Error(err))
	//	return
	//}
	//
	//// init logger
	//if err = srlogger.InitLogger(conf.Cfg.LogConfig, "dev"); err != nil {
	//	zap.L().Warn("error init logger", zap.Error(err))
	//	return
	//}
	//// init postgresql
	//if err = postgresql.Init(conf.Cfg.PostgresConfig); err != nil {
	//	zap.L().Warn("error init postgres", zap.Error(err))
	//	return
	//}
	//// init snowflake
	//if err := snowflake.Init(conf.Cfg.MachineId); err != nil {
	//	zap.L().Error("error init snowflake", zap.Error(err))
	//	return
	//}

	// add router
	router.SetUpRouter(r, conf.Cfg.Mode)

	srlogger.Logger().Info("init server success")
}

func (s *Server) Start() {
	// run srserver
	var err error
	if err = s.Engine.Run(fmt.Sprintf(":%d", conf.Cfg.Port)); err != nil {
		zap.L().Warn("error run srserver", zap.Error(err))
		return
	}
	srlogger.Logger().Info("srserver start success")
}
