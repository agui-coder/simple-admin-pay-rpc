package order

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-common/consts"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderExtensionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderExtensionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderExtensionLogic {
	return &CreateOrderExtensionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderExtensionLogic) CreateOrderExtension(in *pay.OrderCreateExtensionReq) (*pay.BaseIDResp, error) {
	orderExtension, err := l.svcCtx.DB.OrderExtension.Create().
		SetOrderID(in.OrderID).
		SetChannelCode(in.ChannelCode).
		SetChannelExtras(in.ChannelExtras).
		SetNo(in.No).
		SetChannelID(in.ChannelID).
		SetStatus(consts.WAITING).SetUserIP(in.UserIP).Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.BaseIDResp{Id: orderExtension.ID, Msg: i18n.CreateSuccess}, nil
}
