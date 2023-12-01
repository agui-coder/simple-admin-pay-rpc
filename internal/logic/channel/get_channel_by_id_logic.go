package channel

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChannelByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChannelByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChannelByIdLogic {
	return &GetChannelByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetChannelByIdLogic) GetChannelById(in *pay.IDReq) (*pay.ChannelInfo, error) {
	channel, err := l.svcCtx.DB.Channel.Get(l.ctx, in.Id)
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
