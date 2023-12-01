package channel

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/ent/channel"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChannelListByAppIdAndCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChannelListByAppIdAndCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChannelListByAppIdAndCodeLogic {
	return &GetChannelListByAppIdAndCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetChannelListByAppIdAndCodeLogic) GetChannelListByAppIdAndCode(in *pay.ByAppIdAndCodeReq) (*pay.ChannelInfo, error) {
	channel, err := l.svcCtx.DB.Channel.Query().Where(channel.AppIDEQ(in.AppId), channel.CodeEQ(in.Code)).Only(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.ChannelInfo{
		Id:        &channel.ID,
		CreatedAt: pointy.GetPointer(channel.CreatedAt.UnixMilli()),
		UpdatedAt: pointy.GetPointer(channel.UpdatedAt.UnixMilli()),
		Status:    pointy.GetPointer(uint32(channel.Status)),
		Code:      &channel.Code,
		Remark:    &channel.Remark,
		FeeRate:   &channel.FeeRate,
		AppId:     &channel.AppID,
		Config:    &channel.Config,
	}, nil
}
