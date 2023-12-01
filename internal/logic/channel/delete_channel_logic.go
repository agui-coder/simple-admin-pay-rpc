package channel

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteChannelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteChannelLogic {
	return &DeleteChannelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteChannelLogic) DeleteChannel(in *pay.IDReq) (*pay.BaseResp, error) {
	err := l.svcCtx.DB.Channel.DeleteOneID(in.Id).Exec(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.NewInvalidArgumentError("app not found")
		}
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
