package svc

import (
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/config"
	"github.com/agui-coder/simple-admin-pay-rpc/payment"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	//需要导入runtime
	_ "github.com/agui-coder/simple-admin-pay-rpc/ent/runtime"
)

type ServiceContext struct {
	Config      config.Config
	DB          *ent.Client
	Redis       redis.UniversalClient
	AsynqClient *asynq.Client
	PayClient   map[string]model.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()), // driver
		ent.Debug(), // debug mode
	)
	var payConfig payment.PayConfig
	err := c.AliPayConfig.NewAliPayPayConfig(&payConfig)
	if err != nil {
		logx.Error(err.Error())
	}
	err = c.WxPayConfig.NewWxPayPayConfig(&payConfig)
	if err != nil {
		logx.Error(err.Error())
	}
	payClient, err := payment.NewClient(payConfig)
	if err != nil {
		logx.Error(err.Error())
	}
	return &ServiceContext{
		Config:      c,
		DB:          db,
		AsynqClient: c.AsynqConf.WithOriginalRedisConf(c.RedisConf).NewClient(),
		Redis:       c.RedisConf.MustNewUniversalRedis(),
		PayClient:   payClient,
	}
}
