package util_test

import (
	"go.uber.org/zap"
	"srbbs/src/conf"
	"srbbs/src/util/lib/algo"
	"srbbs/src/util/snowflake"
	"testing"
	"time"
)

func TestGetSnowFlake(t *testing.T) {
	var err error
	if err = conf.InitCfg(); err != nil {
		zap.L().Warn("error reading config", zap.Error(err))
		return
	}
	if err := snowflake.Init(conf.Cfg.MachineId); err != nil {
		zap.L().Error("error init snowflake", zap.Error(err))
		return
	}
	id, _ := algo.GetID()
	t.Log("uid: ", id)
	id, _ = algo.GetID()
	t.Log("uid: ", id)

	time.Sleep(1 * time.Second)
}
