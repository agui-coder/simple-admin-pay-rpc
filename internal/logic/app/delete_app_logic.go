package app

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteAppLogic {
	return &DeleteAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteAppLogic) DeleteApp(in *pay.IDReq) (*pay.BaseResp, error) {
	exist, err := l.svcCtx.DB.Order.Query().Where(order.AppIDEQ(in.Id)).Exist(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	if exist {
		return nil, errorx.NewInvalidArgumentError("app exist order cant delete")
	}
	exist, err = l.svcCtx.DB.Refund.Query().Where(refund.AppIDEQ(in.Id)).Exist(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	if exist {
		return nil, errorx.NewInvalidArgumentError("app exist refund cant delete")
	}
	err = l.svcCtx.DB.App.DeleteOneID(in.Id).Exec(l.ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.NewInvalidArgumentError("app not found")
		}
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
