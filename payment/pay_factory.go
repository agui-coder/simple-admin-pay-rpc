package payment

import (
	"fmt"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/ali"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/weixin"
	"github.com/zeromicro/go-zero/core/errorx"
)

type Factory struct {
	clients map[string]model.Client
}

// NewFactory 工厂方法创建支付客户端
func NewFactory() *Factory {
	factory := &Factory{
		clients: make(map[string]model.Client),
	}
	return factory
}

// GetClient 获取支付客户端
func (f *Factory) GetClient(channelCode string) (model.Client, error) {
	client, ok := f.clients[channelCode]
	if !ok {
		return nil, errorx.NewInvalidArgumentError("invalid client ID")
	}
	return client, nil
}

func (f *Factory) ClearClient(channelCode string) {
	delete(f.clients, channelCode)
}

func (f *Factory) CreateOrUpdatePayClient(channelCode string, config model.ClientConfig) error {
	client, exit := f.clients[channelCode]
	if !exit {
		newClient, err := f.getNewClient(channelCode, config)
		if err != nil {
			return err
		}
		err = newClient.Init()
		if err != nil {
			return err
		}
		f.clients[channelCode] = newClient
	} else {
		err := client.Refresh(config)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *Factory) getNewClient(channelCode string, config model.ClientConfig) (model.Client, error) {
	switch channelCode {
	case model.WxApp, model.WxPub, model.WxNative, model.WxLite:
		client := weixin.NewClient(config)
		return client, nil
	case model.AlipayBar, model.AlipayPc, model.AlipayWap, model.AlipayQr:
		client := ali.NewClient(config)
		return client, nil
	}
	return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("unsupported channel code: %s", channelCode))
}

type PayConfig struct {
	AliConfig ali.ClientConfig
	WxConfig  weixin.ClientConfig
}

func NewClient(c PayConfig) (payClient map[string]model.Client, err error) {
	payClient = make(map[string]model.Client)
	if c.AliConfig.Status {
		if err := c.AliConfig.Validate(); err == nil {
			aliClient := ali.NewClient(c.AliConfig)
			err := aliClient.Init()
			if err != nil {
				return nil, err
			}
			payClient[model.AlipayPc] = aliClient
			payClient[model.AlipayBar] = aliClient
			payClient[model.AlipayWap] = aliClient
			payClient[model.AlipayQr] = aliClient
		}
	}
	if c.WxConfig.Status {
		if err := c.WxConfig.Validate(); err == nil {
			wxClient := weixin.NewClient(c.WxConfig)
			err := wxClient.Init()
			if err != nil {
				return nil, err
			}
			payClient[model.WxPub] = wxClient
			payClient[model.WxApp] = wxClient
			payClient[model.WxNative] = wxClient
			payClient[model.WxLite] = wxClient
		}
	}
	return payClient, nil
}
