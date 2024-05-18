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

	// 将 Elasticsearch 客户端注入到 Gin 的 Context 中
	//r.Use(func(c *gin.Context) {
	//	c.Set("esClient", middleware.GetElasticsearchClient())
	//	c.Next()
	//})

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
