package channel

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateChannelCanSubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateChannelCanSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateChannelCanSubmitLogic {
	return &ValidateChannelCanSubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateChannelCanSubmitLogic) ValidateChannelCanSubmit(in *pay.ValidateChannelReq) (*pay.ChannelInfo, error) {
	_, err := l.svcCtx.Model.App.ValidPayApp(l.ctx, in.AppId)
	if err != nil {
		return nil, err
	}
	channel, err := l.svcCtx.Model.Channel.ValidPayChannelByAppIdAndCode(l.ctx, in.AppId, in.Code)
	if err != nil {
		return nil, err
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
