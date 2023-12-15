package svc

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-common/payment"
	"github.com/agui-coder/simple-admin-pay-common/payment/model"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/config"
	entModel "github.com/agui-coder/simple-admin-pay-rpc/model"
	"github.com/zeromicro/go-zero/core/collection"
	"log"
	"strconv"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"

	//需要导入runtime
	_ "github.com/agui-coder/simple-admin-pay-rpc/ent/runtime"
)

type ServiceContext struct {
	Config           config.Config
	DB               *ent.Client
	Model            *entModel.Model
	Redis            *redis.Redis
	AsynqClient      *asynq.Client
	PayClientFactory *payment.Factory
	PayClientCache   *collection.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewCacheDriver(c.RedisConf)), // driver
		ent.Debug(), // debug mode
	)
	cache, err := collection.NewCache(time.Second*10, collection.WithName("payClientCache"))
	if err != nil {
		log.Fatal(err)
	}
	return &ServiceContext{
		Config:           c,
		DB:               db,
		Model:            entModel.NewModel(db),
		AsynqClient:      c.AsynqConf.WithRedisConf(c.RedisConf).NewClient(),
		Redis:            redis.MustNewRedis(c.RedisConf),
		PayClientFactory: payment.NewFactory(),
		PayClientCache:   cache,
	}
}

func (s *ServiceContext) GetPayClient(ctx context.Context, id uint64) (model.Client, error) {
	// 每10更新一次支付客户端
	take, err := s.PayClientCache.Take("pay_client:"+strconv.FormatUint(id, 10), func() (any, error) {
		channel, err := s.Model.Channel.Get(ctx, id)
		if err == nil {
			payConfig, err := payment.ParseClientConfig(channel.Code, channel.Config)
			if err != nil {
				return nil, err
			}
			err = s.PayClientFactory.CreateOrUpdatePayClient(channel.ID, channel.Code, payConfig)
			if err != nil {
				return nil, err
			}
		}
		client, err := s.PayClientFactory.GetClient(channel.ID)
		if err != nil {
			return nil, err
		}
		return client, nil
	})
	if err != nil {
		return nil, err
	}
	return take.(model.Client), nil
}
