package model

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/consts"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"

	"github.com/agui-coder/simple-admin-pay-rpc/ent/channel"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChannelModel struct {
	*ent.ChannelClient
}

func NewChannelModel(client *ent.Client) *ChannelModel {
	return &ChannelModel{ChannelClient: client.Channel}
}

func (m *ChannelModel) ValidPayChannelByAppIdAndCode(ctx context.Context, appId uint64, code string) (*ent.Channel, error) {
	channel, err := m.Query().Where(channel.AppIDEQ(appId), channel.CodeEQ(code)).Only(ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, appId)
	}
	if consts.Disable == channel.Status {
		return nil, errorx.NewInvalidArgumentError("channel is disable")
	}
	return channel, nil
}
