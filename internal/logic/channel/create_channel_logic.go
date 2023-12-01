package channel

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateChannelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChannelLogic {
	return &CreateChannelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Channel management
func (l *CreateChannelLogic) CreateChannel(in *pay.ChannelCreateReq) (*pay.BaseIDResp, error) {
	channel, err := l.svcCtx.Model.Channel.Create().
		SetCode(in.Code).
		SetConfig(in.Config).
		SetNotNilStatus(pointy.GetStatusPointer(&in.Status)).
		SetNotNilRemark(in.Remark).
		SetFeeRate(in.FeeRate).SetAppID(in.AppId).Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.BaseIDResp{Id: channel.ID, Msg: i18n.CreateSuccess}, nil
}
