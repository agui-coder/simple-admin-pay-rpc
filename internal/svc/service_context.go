package svc

import (
	"github.com/agui-coder/simple-admin-pay-rpc/internal/config"
	entModel "github.com/agui-coder/simple-admin-pay-rpc/model"

	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"

	//需要导入runtime
	_ "github.com/agui-coder/simple-admin-pay-rpc/ent/runtime"
)

type ServiceContext struct {
	Config      config.Config
	DB          *ent.Client
	Model       *entModel.Model
	Redis       *redis.Redis
	AsynqServer *asynq.Server
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)
	return &ServiceContext{
		Config:      c,
		DB:          db,
		Model:       entModel.NewModel(db),
		AsynqServer: c.AsynqConf.WithRedisConf(c.RedisConf).NewServer(),
		Redis:       redis.MustNewRedis(c.RedisConf),
	}
}
