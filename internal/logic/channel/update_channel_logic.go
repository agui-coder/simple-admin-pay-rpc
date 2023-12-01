package channel

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateChannelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateChannelLogic {
	return &UpdateChannelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateChannelLogic) UpdateChannel(in *pay.ChannelUpdateReq) (*pay.BaseResp, error) {
	err := l.svcCtx.DB.Channel.UpdateOneID(in.Id).
		SetCode(in.Code).
		SetConfig(in.Config).
		SetNotNilStatus(pointy.GetStatusPointer(&in.Status)).
		SetNotNilRemark(in.Remark).
		SetFeeRate(in.FeeRate).SetAppID(in.AppId).Exec(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.NewInvalidArgumentError("app not found")
		}
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
